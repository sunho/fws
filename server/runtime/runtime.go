package runtime

import (
	"io"

	"github.com/sunho/fws/server/model"
)

type BuildCallback func(err error, result string, logged []byte)

type Builder interface {
	Build(bot *model.Bot, cb BuildCallback) (Building, error)
}

type Building interface {
	Stop() error
	Step() string
}

type Runner interface {
	Exists(bot *model.RunBot) (bool, error)
	Put(bot *model.RunBot) error
	Delete(bot *model.Bot) error

	Run(bot *model.Bot) error
	Restart(bot *model.Bot) error
	Stop(bot *model.Bot) error
	Status(bot *model.Bot) (model.RunStatus, error)
	Log(bot *model.Bot) ([]byte, error)

	UpdateBuild(bot *model.Bot) error
	DownloadVolume(volume *model.BotVolume) (io.Reader, error)
	VolumeUsed(volume *model.BotVolume) (int64, error)
}
