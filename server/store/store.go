package store

import (
	"errors"

	"github.com/sunho/fws/server/model"
)

var (
	ErrNoEntry = errors.New("store: no such entry")
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
	GetBot(id int) (*model.Bot, error)
	CreateBot(bot *model.Bot) (*model.Bot, error)
	UpdateBot(bot *model.Bot) error
	DeleteBot(bot *model.Bot) error

	ListBotConfig(id int) ([]*model.Config, error)
	CreateBotConfig(config *model.Config) (*model.Config, error)
	UpdateBotConfig(config *model.Config) error
	DeleteBotConfig(config *model.Config) error

	ListBotVolume(id int) ([]*model.Volume, error)
	CreateBotVolume(volume *model.Volume) (*model.Volume, error)
	UpdateBotVolume(volume *model.Volume) error
	DeleteBotVolume(volume *model.Volume) error

	ListBotEnv(id int) ([]*model.Env, error)
	CreateBotEnv(env *model.Env) (*model.Env, error)
	UpdateBotEnv(env *model.Env) error
	DeleteBotEnv(env *model.Env) error

	ListUserBot(user int) ([]*model.Bot, error)
	GetUserBot(user int, bot int) (bool, error)
	CreateUserBot(user int, bot int) error
	DeleteUserBot(user int, bot int) error

	ListBotBuild(bot int) ([]*model.Build, error)
	CreateBotBuild(build *model.Build) (*model.Build, error)
	DeleteBotBuild(build *model.Build) error

	GetBotBuildLog(bot int, number int) (*model.BuildLog, error)
	CreateBotBuildLog(build *model.BuildLog) (*model.BuildLog, error)
}
