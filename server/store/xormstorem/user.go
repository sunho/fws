package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) GetUser(id int) (*model.User, error) {
	u := &model.User{
		ID: id,
	}
	has, err := x.e.Get(u)
	if !has {
		return nil, store.ErrNoEntry
	}
	return u, err
}

func (x *XormStore) GetUserByUsername(username string) (*model.User, error) {
	u := &model.User{
		Username: username,
	}
	has, err := x.e.Get(u)
	if !has {
		return nil, store.ErrNoEntry
	}
	return u, err
}

func (x *XormStore) GetUserByNickname(nickname string) (*model.User, error) {
	u := &model.User{
		Nickname: nickname,
	}
	has, err := x.e.Get(u)
	if !has {
		return nil, store.ErrNoEntry
	}
	return u, err
}

func (x *XormStore) CreateUser(user *model.User) (*model.User, error) {
	_, err := x.e.Insert(user)
	return user, err
}

func (x *XormStore) UpdateUser(user *model.User) error {
	_, err := x.e.Update(user)
	return err
}

func (x *XormStore) DeleteUser(user *model.User) error {
	_, err := x.e.ID(user.ID).Delete(new(model.User))
	return err
}
