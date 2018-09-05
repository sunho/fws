package migrations

import (
	"errors"

	"github.com/go-xorm/xorm"
)

const CurrentV = 0

var (
	ErrSuperior = errors.New("xormstore/migrations: no forward compatibility")
)

type Version struct {
	ID      int64 `xorm:"pk autoincr"`
	Version int64
}

func Migrate(e *xorm.Engine) error {
	err := e.Sync(new(Version))
	if err != nil {
		return err
	}

	v := &Version{
		ID: 1,
	}

	has, err := e.Get(v)
	if err != nil {
		return err
	} else if !has {
		v.ID = 0
		v.Version = CurrentV

		if _, err = e.InsertOne(v); err != nil {
			return err
		}
		return nil
	}

	if v.Version > CurrentV {
		return ErrSuperior
	} else if v.Version == CurrentV {
		return nil
	} else {
		panic("to do")
	}
}
