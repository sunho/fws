package store

import (
	"errors"

	"github.com/sunho/fws/server/model"
)

var (
	ErrNotExists = errors.New("store: doesn't exist")
)

type Store interface {
	ListUser() ([]*model.User, error)
	GetUser(id int) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByNickname(nickname string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error

	ListUserInvite() ([]*model.UserInvite, error)
	GetUserInvite(username string) (*model.UserInvite, error)
	CreateUserInvite(i *model.UserInvite) (*model.UserInvite, error)
	DeleteUserInvite(i *model.UserInvite) error

	ListBot() ([]*model.Bot, error)
	GetBot(bot int) (*model.Bot, error)
	CreateBot(bot *model.Bot) (*model.Bot, error)
	UpdateBot(bot *model.Bot) error
	DeleteBot(bot *model.Bot) error

	ListBotConfig(bot int) ([]*model.BotConfig, error)
	GetBotConfig(bot int, name string) (*model.BotConfig, error)
	CreateBotConfig(config *model.BotConfig) (*model.BotConfig, error)
	UpdateBotConfig(config *model.BotConfig) error
	DeleteBotConfig(config *model.BotConfig) error

	ListBotVolume(bot int) ([]*model.BotVolume, error)
	GetBotVolume(bot int, vol string) (*model.BotVolume, error)
	CreateBotVolume(volume *model.BotVolume) (*model.BotVolume, error)
	UpdateBotVolume(volume *model.BotVolume) error
	DeleteBotVolume(volume *model.BotVolume) error

	ListBotEnv(bot int) ([]*model.BotEnv, error)
	GetBotEnv(bot int, env string) (*model.BotEnv, error)
	CreateBotEnv(env *model.BotEnv) (*model.BotEnv, error)
	UpdateBotEnv(env *model.BotEnv) error
	DeleteBotEnv(env *model.BotEnv) error

	ListUserBot(user int) ([]*model.Bot, error)
	GetUserBot(user int, bot int) (bool, error)
	CreateUserBot(user int, bot int) error
	DeleteUserBot(user int, bot int) error

	ListBotBuild(bot int) ([]*model.Build, error)
	CreateBotBuild(build *model.Build) (*model.Build, error)
	DeleteBotBuild(build *model.Build) error

	GetBotBuildLog(bot int, number int) (*model.BuildLog, error)
	CreateBotBuildLog(build *model.BuildLog) (*model.BuildLog, error)

	ListBotRun(bot int) ([]*model.Run, error)
	GetLatestBotRun(bot int) (*model.Run, error)
	CreateBotRun(l *model.Run) (*model.Run, error)
	UpdateBotRun(l *model.Run) error
	DeleteBotRun(l *model.Run) error

	GetBotRunLog(bot int, number int) (*model.RunLog, error)
	CreateBotRunLog(run *model.RunLog) (*model.RunLog, error)
	UpdateBotRunLog(run *model.RunLog) error
}
