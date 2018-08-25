package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) GetUserInvite(username string) (*model.UserInvite, error) {
	i := &model.UserInvite{
		Username: username,
	}
	has, err := x.e.Get(i)
	if !has {
		return nil, store.ErrNoEntry
	}
	return i, err
}

func (x *XormStore) CreateUserInvite(i *model.UserInvite) (*model.UserInvite, error) {
	_, err := x.e.Insert(i)
	return i, err
}

func (x *XormStore) DeleteUserInvite(i *model.UserInvite) error {
	_, err := x.e.Where("username = ?", i.Username).
		Delete(new(model.UserInvite))
	return err
}
