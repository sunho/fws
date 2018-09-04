package fws

import (
	"sync"

	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

const maxCurrent = 10

type BuildManager struct {
	mu *sync.RWMutex

	check   chan struct{}
	current int
	builds  map[int]*builds
}

type build struct {
	bot *model.Bot
}

func (b *BuildManager) Start() {
	go func() {
		for {
			select {
			case <-b.check:
				b.startPendingBuilds()
			}
		}
	}()
}

func (b *BuildManager) Build(bot *model.Bot) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.buildings[bot.ID]; ok {
		return runtime.ErrAlreadyBuilding
	}

	cb2 := func(err error, logged []byte) {
		b.mu.Lock()
		b.current--
		delete(b.buildings, bot.ID)
		b.mu.Unlock()

		cb(err, logged)
	}

	b.buildings[bot.ID] = &building{
		parent: b,
		bot:    bot,
		cb:     cb2,
		logged: []byte{},
	}
	b.check <- struct{}{}

	return nil
}

func (b *BuildManager) startPendingBuilds() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, bui := range b.buildings {
		if b.current >= maxCurrent {
			return
		}
		if !bui.running {
			b.current++
			bui.Start()
		}
	}
}

func (b *BuildManager) Stop(bot *model.Bot) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bui, ok := b.buildings[bot.ID]
	if !ok {
		return runtime.ErrNotExists
	}
	return bui.Stop()
}

func (b *BuildManager) Status(bot *model.Bot) (model.BuildStatus, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bui, ok := b.buildings[bot.ID]
	if !ok {
		return model.BuildStatus{}, runtime.ErrNotExists
	}
	return bui.Status(), nil
}
