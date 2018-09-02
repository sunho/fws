package basic

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

const maxCurrent = 10

type DefaultBuilder struct {
	mu        *sync.RWMutex
	check     chan struct{}
	current   int
	buildings map[int]*building

	RegURL    string
	Workspace string
}

func (b *DefaultBuilder) Start() {
	go func() {
		for {
			select {
			case <-b.check:
			}
		}
	}()
}

func (b *DefaultBuilder) Build(bot *model.Bot, cb runtime.BuildCallback) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.buildings[bot.ID]; ok {
		return runtime.ErrAlreadyBuilding
	}

	b.buildings[bot.ID] = &building{
		parent: b,
		bot:    bot,
		cb:     cb,
		logged: []byte{},
	}
	b.check <- struct{}{}

	return nil
}

func (b *DefaultBuilder) Stop(bot *model.Bot) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bui, ok := b.buildings[bot.ID]
	if !ok {
		return runtime.ErrNotExists
	}
	return bui.Stop()
}

func (b *DefaultBuilder) Status(bot *model.Bot) (model.BuildStatus, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bui, ok := b.buildings[bot.ID]
	if !ok {
		return model.BuildStatus{}, runtime.ErrNotExists
	}
	return bui.Status(), nil
}

type building struct {
	mu     *sync.RWMutex
	parent *DefaultBuilder
	bot    *model.Bot
	cb     runtime.BuildCallback
	kill   func()

	running bool
	step    string
	logged  []byte
}

func (b *building) Start() {
	b.running = true
	go func() {
		err := b.work()
		b.cb(err)
	}()
}

func (b *building) Stop() error {
	//TODO
	b.kill()
	return nil
}

func (b *building) Status() model.BuildStatus {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return model.BuildStatus{
		Running: b.running,
		Step:    b.step,
	}
}

func (b *building) work() error {
	path := b.parent.Workspace + "/" + strconv.Itoa(b.bot.ID)
	img := fmt.Sprintf("%s/%s%d:%d", b.parent.RegURL, b.bot.Name, b.bot.ID, time.Now().Unix())

	b.setStep("clean")
	rm := exec.Command("rm", "-rf", path)
	if err := b.exec("clean", rm); err != nil {
		return err
	}

	b.setStep("download")
	clone := exec.Command("git", "clone", b.bot.GitURL, path)
	if err := b.exec("download", clone); err != nil {
		return err
	}

	b.setStep("build")
	build := exec.Command("docker", "build", "-t", img)
	build.Dir = path
	if err := b.exec("build", build); err != nil {
		return err
	}

	b.setStep("upload")
	push := exec.Command("docker", "push", img)
	push.Dir = path
	if err := b.exec("upload", push); err != nil {
		return err
	}

	return nil
}

func (b *building) exec(name string, cmd *exec.Cmd) error {
	b.kill = func() {
		// TODO
		cmd.Process.Kill()
	}

	r, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	r2, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	b.writeLog([]byte("-----" + name + "-----"))

	var wg sync.WaitGroup
	go b.streamLog(&wg, r)
	wg.Add(1)
	go b.streamLog(&wg, r2)
	wg.Add(1)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func (b *building) streamLog(wg *sync.WaitGroup, r io.Reader) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		buf := s.Text()
		b.writeLog([]byte(buf))
	}
	wg.Done()
}

func (b *building) setStep(str string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.step = str
}

func (b *building) writeLog(buf []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logged = append(b.logged, buf...)
}
