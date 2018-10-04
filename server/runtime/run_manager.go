package runtime

import (
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

type RunManager struct {
	mu sync.RWMutex

	runner Runner
	stor   store.Store
}

func NewRunManager(runner Runner, stor store.Store) *RunManager {
	return &RunManager{
		runner: runner,
		stor:   stor,
	}
}

func (r *RunManager) Start() {
	go func() {
		t := time.NewTicker(time.Minute)
		for {
			select {
			case <-t.C:
				r.collectLog()
				r.cleanLog()
			}
		}
	}()
}

func (r *RunManager) createRun(b *model.Bot, until time.Time) {
	nrun, err := r.stor.CreateBotRun(&model.Run{
		BotID: b.ID,
		Since: time.Now(),
		Until: until,
	})
	if err != nil {
		glog.Errorf("Error while creating BotRun, err: %v", err)
		return
	}
	_, err = r.stor.CreateBotRunLog(&model.RunLog{
		BotID:  b.ID,
		Number: nrun.Number,
		Logged: []byte{},
	})
	if err != nil {
		glog.Errorf("Error while creating BotRunLog, err: %v", err)
		return
	}
}

func (r *RunManager) cleanLog() {
	bots, err := r.stor.ListBot()
	if err != nil {
		glog.Errorf("RunManager.cleanLog faild, err: %v", err)
		return
	}
	for _, b := range bots {
		runs, err := r.stor.ListBotRun(b.ID)
		if err != nil {
			glog.Errorf("Error while listing BotRun, err: %v", err)
			continue
		}
		for _, ru := range runs {
			if time.Now().Sub(ru.Since) >= time.Hour*24*4 {
				err = r.stor.DeleteBotRun(ru)
				if err != nil {
					glog.Errorf("Error while deleting BotRun, err: %v", err)
					continue
				}
			}
		}
	}
}

func (r *RunManager) collectLog() {
	bots, err := r.stor.ListBot()
	if err != nil {
		glog.Errorf("RunManager.collectLog faild, err: %v", err)
		return
	}

	for _, b := range bots {
		run, err := r.stor.GetLatestBotRun(b.ID)
		if err == store.ErrNotExists {
			r.createRun(b, time.Now())
			continue
		} else if err != nil {
			glog.Errorf("Error while getting BotRun, err: %v", err)
			continue
		}

		if time.Now().Sub(run.Since) >= 24*time.Hour {
			r.createRun(b, run.Until)
			continue
		}

		buf, _ := r.runner.Log(b, run.Until)
		if buf == nil {
			buf = []byte{}
		}

		rlog, err := r.stor.GetBotRunLog(b.ID, run.Number)
		if err != nil {
			glog.Errorf("Error while getting BotRunLog, err: %v", err)
			continue
		}

		run.Until = time.Now()
		rlog.Logged = append(rlog.Logged, buf...)

		err = r.stor.UpdateBotRun(run)
		if err != nil {
			glog.Errorf("Error while updating BotRun, err: %v", err)
			continue
		}

		err = r.stor.UpdateBotRunLog(rlog)
		if err != nil {
			glog.Errorf("Error while updating BotRunLog, err: %v", err)
			continue
		}
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
