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
func (r *RunManager) Create(bot *model.RunBot) error {
	return r.runner.Create(bot)
}
