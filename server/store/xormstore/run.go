package xormstore

import "github.com/sunho/fws/server/model"

func (x *XormStore) ListBotConfig(bot int) ([]*model.Config, error) {
	var cs []*model.Config
	err := x.e.Where("bot_id = ?", bot).Find(&cs)
	return cs, err
}

func (x *XormStore) CreateBotConfig(config *model.Config) (*model.Config, error) {
	_, err := x.e.Insert(config)
	return config, err
}

func (x *XormStore) UpdateBotConfig(config *model.Config) error {
	_, err := x.e.Update(config)
	return err
}

func (x *XormStore) DeleteBotConfig(config *model.Config) error {
	_, err := x.e.Where("bot_id = ? AND name = ?", config.BotID, config.Name).
		Delete(&model.Config{})
	return err
}

func (x *XormStore) ListBotVolume(bot int) ([]*model.Volume, error) {
	var vs []*model.Volume
	err := x.e.Where("bot_id = ?", bot).Find(&vs)
	return vs, err
}

func (x *XormStore) CreateBotVolume(volume *model.Volume) (*model.Volume, error) {
	_, err := x.e.Insert(volume)
	return volume, err
}

func (x *XormStore) UpdateBotVolume(volume *model.Volume) error {
	_, err := x.e.Where("bot_id = ? AND name = ?", volume.BotID, volume.Name).
		Delete(&model.Volume{})
	return err
}

func (x *XormStore) DeleteBotVolume(volume *model.Volume) error {
	_, err := x.e.Update(volume)
	return err
}

func (x *XormStore) ListBotEnv(bot int) ([]*model.Env, error) {
	var es []*model.Env
	err := x.e.Where("bot_id = ?", bot).Find(&es)
	return es, err
}

func (x *XormStore) CreateBotEnv(env *model.Env) (*model.Env, error) {
	_, err := x.e.Insert(env)
	return env, err
}

func (x *XormStore) UpdateBotEnv(env *model.Env) error {
	_, err := x.e.Update(env)
	return err
}

func (x *XormStore) DeleteBotEnv(env *model.Env) error {
	_, err := x.e.Where("bot_id = ? AND name = ?", env.BotID, env.Name).
		Delete(&model.Env{})
	return err
}
