package runtime

import (
	"sync"

	"github.com/sunho/fws/server/model"
)

type RunManager struct {
	mu sync.RWMutex

	runner Runner
}

func NewRunManager(runner Runner) *RunManager {
	return &RunManager{
		runner: runner,
	}
}

// TODO queue
func (r *RunManager) Put(bot *model.RunBot) error {
	return r.runner.Put(bot)
}

func (r *RunManager) Restart(bot *model.Bot) error {
	return r.runner.Restart(bot)
}

func (r *RunManager) Status(bot *model.Bot) (model.RunStatus, error) {
	return r.runner.Status(bot)
}
