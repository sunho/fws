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
		return nil, store.ErrNotExists
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
	_, err := x.e.Where("bot_id = ?", bot.ID).Delete(new(model.UserBot))
	if err != nil {
		return err
	}

	_, err = x.e.Where("bot_id = ?", bot.ID).Delete(new(model.BotConfig))
	if err != nil {
		return err
	}

	_, err = x.e.Where("bot_id = ?", bot.ID).Delete(new(model.BotEnv))
	if err != nil {
		return err
	}

	_, err = x.e.Where("bot_id = ?", bot.ID).Delete(new(model.BotVolume))
	if err != nil {
		return err
	}

	_, err = x.e.Where("bot_id = ?", bot.ID).Delete(new(model.Build))
	if err != nil {
		return err
	}

	_, err = x.e.Where("bot_id = ?", bot.ID).Delete(new(model.BuildLog))
	if err != nil {
		return err
	}

	_, err = x.e.ID(bot.ID).Delete(&model.Bot{})
	return err
}

func (x *XormStore) ListBotConfig(bot int) ([]*model.BotConfig, error) {
	var cs []*model.BotConfig
	err := x.e.Where("bot_id = ?", bot).Find(&cs)
	return cs, err
}

func (x *XormStore) CreateBotConfig(config *model.BotConfig) (*model.BotConfig, error) {
	_, err := x.e.Insert(config)
	return config, err
}

func (x *XormStore) UpdateBotConfig(config *model.BotConfig) error {
	_, err := x.e.Update(config)
	return err
}

func (x *XormStore) DeleteBotConfig(config *model.BotConfig) error {
	_, err := x.e.Where("bot_id = ? AND name = ?", config.BotID, config.Name).
		Delete(&model.BotConfig{})
	return err
}

func (x *XormStore) ListBotVolume(bot int) ([]*model.BotVolume, error) {
	var vs []*model.BotVolume
	err := x.e.Where("bot_id = ?", bot).Find(&vs)
	return vs, err
}

func (x *XormStore) CreateBotVolume(volume *model.BotVolume) (*model.BotVolume, error) {
	_, err := x.e.Insert(volume)
	return volume, err
}

func (x *XormStore) UpdateBotVolume(volume *model.BotVolume) error {
	_, err := x.e.Where("bot_id = ? AND name = ?", volume.BotID, volume.Name).
		Delete(&model.BotVolume{})
	return err
}

func (x *XormStore) DeleteBotVolume(volume *model.BotVolume) error {
	_, err := x.e.Update(volume)
	return err
}

func (x *XormStore) ListBotEnv(bot int) ([]*model.BotEnv, error) {
	var es []*model.BotEnv
	err := x.e.Where("bot_id = ?", bot).Find(&es)
	return es, err
}

func (x *XormStore) CreateBotEnv(env *model.BotEnv) (*model.BotEnv, error) {
	_, err := x.e.Insert(env)
	return env, err
}

func (x *XormStore) UpdateBotEnv(env *model.BotEnv) error {
	_, err := x.e.Update(env)
	return err
}

func (x *XormStore) DeleteBotEnv(env *model.BotEnv) error {
	_, err := x.e.Where("bot_id = ? AND name = ?", env.BotID, env.Name).
		Delete(&model.BotEnv{})
	return err
}
