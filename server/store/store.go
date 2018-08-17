package store

import "github.com/sunho/bot-registry/server/model"

type Store interface {
	GetBot(id int) (*model.Bot, error)
	CreateBot(bot *model.Bot) (*model.Bot, error)
	UpdateBot(bot *model.Bot) error
	DeleteBot(bot *model.Bot) error

	GetUser(id int) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error

	ListUserBot(id int) ([]*model.Bot, error)
	CreateUserBot(id int, botid int) error
	DeleteUserBot(userid int, botid int) error

	ListBotConfig(id int) ([]*model.Config, error)
	CreateBotConfig(id int, config *model.Config) (*model.Config, error)
	UpdateBotConfig(id int, config *model.Config) error
	DeleteBotConfig(id int, config *model.Config) error

	ListBotVolume(id int) ([]*model.Volume, error)
	CreateBotVolume(id int, volume *model.Volume) (*model.Volume, error)
	UpdateBotVolume(id int, volume *model.Volume) error
	DeleteBotVolume(id int, volume *model.Volume) error

	ListBotEnv(id int) ([]*model.Env, error)
	CreateBotEnv(id int, env *model.Env) (*model.Env, error)
	UpdateBotEnv(id int, env *model.Env) error
	DeleteBotEnv(id int, env *model.Env) error

	ListBotBuild(id int) ([]*model.Build, error)
	CreateBotBuild(id int, build *model.Build) (*model.Build, error)
	DeleteBotBuild(id int, build *model.Build) error
	GetBotBuildLog(id int, build *model.Build) (string, error)

	GetWebhook(hash string) (*model.Webhook, error)
	CreateWebhook(hook *model.Webhook) (*model.Webhook, error)
	DeleteWebhook(hook *model.Webhook) error
}
