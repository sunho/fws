package basic

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

type DefaultBuilder struct {
	RegURL    string
	Workspace string
}

func (b *DefaultBuilder) Build(bot *model.Bot, cb runtime.BuildCallback) error {
	path := b.Workspace + "/" + strconv.Itoa(bot.ID)
	img := fmt.Sprintf("%s%d:%d", bot.Name, bot.ID, time.Now().Unix())

	rm := exec.Command("rm", "-rf", path)

	clone := exec.Command("git", "clone", bot.GitURL, path)

	build := exec.Command("docker", "build", "-t", img)
	build.Dir = path

	push := exec.Command("docker", "push", img)
	push.Dir = path

	return nil
}

func (b *DefaultBuilder) Stop(bot *model.Bot) error {

}

func (b *DefaultBuilder) Status(bot *model.Bot) (model.BuildStatus, error) {

}
