package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) ListBot() ([]*model.Bot, error) {
	var bs []*model.Bot
	err := x.e.Find(&bs)
	return bs, err
}

func (x *XormStore) GetBot(id int) (*model.Bot, error) {
	var b model.Bot
	has, err := x.e.ID(id).Get(&b)
	if !has {
		return nil, store.ErrNoEntry
	}
	return &b, err
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
	_, err := x.e.ID(bot.ID).Delete(&model.Bot{})
	return err
}
