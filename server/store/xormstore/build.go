package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) ListBotBuild(bot int) ([]*model.Build, error) {
	var out []*model.Build
	err := x.e.Where("bot_id = ?", bot).Find(&out)
	return out, err
}

func (x *XormStore) CreateBotBuild(build *model.Build) (*model.Build, error) {
	builds, err := x.ListBotBuild(build.BotID)
	if err != nil {
		return nil, err
	}

	// potential bottleneck
	var num int
	for _, build := range builds {
		if build.Number > num {
			num = build.Number
		}
	}
	build.Number = num + 1

	_, err = x.e.Insert(build)
	return build, err
}

func (x *XormStore) DeleteBotBuild(build *model.Build) error {
	_, err := x.e.Where("bot_id = ? AND number = ?", build.BotID, build.Number).
		Delete(new(model.Build))
	return err
}

func (x *XormStore) GetBotBuildLog(bot int, number int) (*model.BuildLog, error) {
	var b *model.BuildLog
	has, err := x.e.Where("bot_id = ? AND number = ?", bot, number).Get(b)
	if !has {
		return nil, store.ErrNoEntry
	}
	return b, err
}

func (x *XormStore) CreateBotBuildLog(build *model.BuildLog) (*model.BuildLog, error) {
	_, err := x.e.Insert(build)
	return build, err
}
