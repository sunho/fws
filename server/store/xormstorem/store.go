package xormstore

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/store/xormstorem/migrations"
)

type XormStore struct {
	e *xorm.Engine
}

func New(e *xorm.Engine) *XormStore {
	e.SetMapper(core.GonicMapper{})
	e.ShowSQL(true)
	return &XormStore{e}
}

func (x *XormStore) Migrate() error {
	err := migrations.Migrate(x.e)
	if err != nil {
		return err
	}

	err = x.e.Sync(new(model.User), new(model.UserInvite))
	if err != nil {
		return err
	}

	return nil
}