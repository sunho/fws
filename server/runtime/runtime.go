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
	CreateBot(bot *model.RunBot) error
	UpdateBot(bot *model.RunBot) error
	DeleteBot(bot *model.Bot) error

	RunBot(bot *model.Bot) error
	RestartBot(bot *model.Bot) error
	StopBot(bot *model.Bot) error
	BotStatus(bot *model.Bot) (model.RunStatus, error)
	BotLog(bot *model.Bot) ([]byte, error)

	DownloadVolume(volume *model.BotVolume) (io.Reader, error)
	VolumeUsed(volume *model.BotVolume) (int64, error)
}
