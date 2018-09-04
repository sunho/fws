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

type Builder struct {
	RegURL    string
	Workspace string
}

func (b *Builder) Build(bot *model.Bot, cb runtime.BuildCallback) (*Building, error) {
	bui := &Building{
		parent: b,
		bot:    bot,
		cb:     cb,
		logged: []byte{},
	}
	bui.Start()
	return bui, nil
}

type Building struct {
	mu     *sync.RWMutex
	parent *Builder
	bot    *model.Bot
	cb     runtime.BuildCallback
	kill   func()

	step   string
	logged []byte
}

func (b *Building) Start() {
	go func() {
		err := b.work()
		b.cb(err, b.logged)
	}()
}

func (b *Building) Stop() error {
	//TODO
	b.kill()
	return nil
}

func (b *Building) Step() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.step
}

func (b *Building) work() error {
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

func (b *Building) exec(name string, cmd *exec.Cmd) error {
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

func (b *Building) streamLog(wg *sync.WaitGroup, r io.Reader) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		buf := s.Text()
		b.writeLog([]byte(buf))
	}
	wg.Done()
}

func (b *Building) setStep(str string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.step = str
}

func (b *Building) writeLog(buf []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logged = append(b.logged, buf...)
}
