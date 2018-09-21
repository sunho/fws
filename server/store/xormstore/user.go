package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) ListUser() ([]*model.User, error) {
	var us []*model.User
	err := x.e.Find(&us)
	return us, err
}

func (x *XormStore) GetUser(id int) (*model.User, error) {
	var u model.User
	has, err := x.e.ID(id).Get(&u)
	if !has {
		return nil, store.ErrNotExists
	}
	return &u, err
}

func (x *XormStore) GetUserByUsername(username string) (*model.User, error) {
	var u model.User
	has, err := x.e.Where("username = ?", username).Get(&u)
	if !has {
		return nil, store.ErrNotExists
	}
	return &u, err
}

func (x *XormStore) GetUserByNickname(nickname string) (*model.User, error) {
	var u model.User
	has, err := x.e.Where("nickname = ?", nickname).Get(&u)
	if !has {
		return nil, store.ErrNotExists
	}
	return &u, err
}

func (x *XormStore) CreateUser(user *model.User) (*model.User, error) {
	_, err := x.e.Insert(user)
	return user, err
}

func (x *XormStore) UpdateUser(user *model.User) error {
	_, err := x.e.Where("id = ?", user.ID).Update(user)
	return err
}

func (x *XormStore) DeleteUser(user *model.User) error {
	_, err := x.e.Where("user_id = ?", user.ID).Delete(new(model.UserBot))
	if err != nil {
		return err
	}

	_, err = x.e.ID(user.ID).Delete(new(model.User))
	if err != nil {
		return err
	}

	return err
}

func (x *XormStore) ListUserBot(user int) ([]*model.Bot, error) {
	var bots []*model.Bot
	err := x.e.Table("user_bot").
		Where("user_id = ?", user).
		Join("INNER", "bot", "user_bot.bot_id = bot.id").
		Find(&bots)
	return bots, err
}

func (x *XormStore) GetUserBot(user int, bot int) (bool, error) {
	var b model.UserBot
	has, err := x.e.Where("user_id = ? AND bot_id = ?", user, bot).
		Get(&b)
	if !has {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (x *XormStore) CreateUserBot(user int, bot int) error {
	_, err := x.e.Insert(&model.UserBot{
		UserID: user,
		BotID:  bot,
	})
	return err
}

func (x *XormStore) DeleteUserBot(user int, bot int) error {
	_, err := x.e.Where("user_id = ? AND bot_id = ?", user, bot).
		Delete(new(model.UserBot))
	return err
}
