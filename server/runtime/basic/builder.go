package basic

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func NewBuilder(regurl string, workspace string) *Builder {
	return &Builder{
		RegURL:    regurl,
		Workspace: workspace,
	}
}

func (b *Builder) Build(bot *model.Bot, cb runtime.BuildCallback) (runtime.Building, error) {
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
	mu     sync.RWMutex
	parent *Builder
	bot    *model.Bot
	cb     runtime.BuildCallback
	kill   func()

	img    string
	step   string
	logged []byte
}

func (b *Building) Start() {
	go func() {
		err := b.work()
		if err != nil {
			b.writeLog([]byte("error:" + err.Error() + "\n"))
		}
		b.cb(err, b.img, b.logged)
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
	b.img = fmt.Sprintf("%s/%s%d:%d", b.parent.RegURL, b.bot.Name, b.bot.ID, time.Now().Unix())

	_, err := os.Stat(b.parent.Workspace)
	if os.IsNotExist(err) {
		err = os.Mkdir(b.parent.Workspace, 0644)
		if err != nil {
			return err
		}
	}

	b.setStep("clean")
	_, err = os.Stat(path)
	if err == nil {
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}

	b.setStep("download")
	clone := exec.Command("git", "clone", b.bot.GitURL, path)
	if err := b.exec("download", clone); err != nil {
		return err
	}

	b.setStep("build")
	build := exec.Command("docker", "build", "-t", b.img)
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

	b.writeLog([]byte("-----" + name + "-----\n"))

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
