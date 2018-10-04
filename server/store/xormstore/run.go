package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) ListBotRun(bot int) ([]*model.Run, error) {
	var rs []*model.Run
	err := x.e.Where("bot_id = ?", bot).Find(&rs)
	return rs, err
}

func (x *XormStore) CreateBotRun(run *model.Run) (*model.Run, error) {
	runs, err := x.ListBotRun(run.BotID)
	if err != nil {
		return nil, err
	}

	var num int
	for _, r := range runs {
		if r.Number > num {
			num = r.Number
		}
	}
	run.Number = num + 1
	_, err = x.e.Insert(run)
	return run, err
}

func (x *XormStore) UpdateBotRun(run *model.Run) error {
	_, err := x.e.Where("bot_id = ? AND number = ?", run.BotID, run.Number).
		Update(run)
	return err
}

func (x *XormStore) DeleteBotRun(run *model.Run) error {
	_, err := x.e.Where("bot_id = ? AND number = ?", run.BotID, run.Number).
		Delete(new(model.Run))
	if err != nil {
		return err
	}
	_, err = x.e.Where("bot_id = ? AND number = ?", run.BotID, run.Number).
		Delete(new(model.RunLog))
	return err
}

func (x *XormStore) GetBotRunLog(bot int, number int) (*model.RunLog, error) {
	var r model.RunLog
	has, err := x.e.Where("bot_id = ? AND number = ?", bot, number).Get(&r)
	if !has {
		return nil, store.ErrNotExists
	}
	return &r, err
}

func (x *XormStore) GetLatestBotRun(bot int) (*model.Run, error) {
	runs, err := x.ListBotRun(bot)
	if err != nil {
		return nil, err
	}

	var temp *model.Run
	for _, r := range runs {
		if temp == nil || r.Since.After(temp.Since) {
			temp = r
		}
	}

	if temp == nil {
		return nil, store.ErrNotExists
	}

	return temp, nil
}

func (x *XormStore) CreateBotRunLog(run *model.RunLog) (*model.RunLog, error) {
	_, err := x.e.Insert(run)
	return run, err
}

func (x *XormStore) UpdateBotRunLog(run *model.RunLog) error {
	_, err := x.e.Where("bot_id = ? AND number = ?", run.BotID, run.Number).
		Update(run)
	return err
}
