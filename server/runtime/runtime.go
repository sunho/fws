package runtime

import (
	"errors"
	"io"

	"github.com/sunho/fws/server/model"
)

var (
	ErrAlreadyBuilding = errors.New("runtime: already building")
	ErrNotExists       = errors.New("runtime: doesn't exists")
)

type BuildCallback func(error)

type Builder interface {
	Build(bot *model.Bot, cb BuildCallback) error
	Stop(bot *model.Bot) error
	Status(bot *model.Bot) (model.BuildStatus, error)
}

type Runner interface {
	CreateBot(bot *model.RunBot) error
	UpdateBot(bot *model.RunBot) error
	DeleteBot(id int) error

	RunBot(id int) error
	RestartBot(id int) error
	StopBot(id int) error
	BotStatus(id int) (model.RunStatus, error)
	BotLog(id int) (string, error)

	DownloadVolume(volume *model.Volume) (io.Reader, error)
	VolumeUsed(volume *model.Volume) (int64, error)
}
