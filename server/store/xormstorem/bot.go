package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) GetBot(id int) (*model.Bot, error) {
	b := &model.Bot{
		ID: id,
	}
	has, err := x.e.Get(b)
	if !has {
		return nil, store.ErrNoEntry
	}
	return b, err
}

func (x *XormStore) CreateBot(bot *model.Bot) (*model.Bot, error) {
	bot.ID = 0
	_, err := x.e.Insert(bot)
	return bot, err
}

func (x *XormStore) UpdateBot(bot *model.Bot) error {
	_, err := x.e.Update(bot)
	return err
}

func (x *XormStore) DeleteBot(bot *model.Bot) error {
	_, err := x.e.Delete(&model.Bot{
		ID: bot.ID,
	})
	return err
}
