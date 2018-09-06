package xormstore

import (
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store"
)

func (x *XormStore) ListUserInvite() ([]*model.UserInvite, error) {
	var us []*model.UserInvite
	err := x.e.Find(&us)
	return us, err
}

func (x *XormStore) GetUserInvite(username string) (*model.UserInvite, error) {
	var i model.UserInvite
	has, err := x.e.Where("username = ?", username).Get(&i)
	if !has {
		return nil, store.ErrNotExists
	}
	return &i, err
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
