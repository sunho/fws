package fws

import (
	"net/http"

	"github.com/sunho/fws/server/runtime"
	"github.com/sunho/fws/server/store"
)

type fwsInterface struct {
	f *Fws
}

func (f *fwsInterface) GetBuilder() runtime.Builder {
	return f.f.builder
}

func (f *fwsInterface) GetRunner() runtime.Runner {
	return f.f.runner
}

func (f *fwsInterface) GetStore() store.Store {
	return f.f.stor
}

func (f *fwsInterface) CreateInviteKey(username string) string {
	return "todo"
}

func (f *fwsInterface) HashPassword(password string) string {
	return "todo"
}

func (f *fwsInterface) CreateToken(id int, username string) string {
	return "todo"
}

func (f *fwsInterface) ParseToken(tok string) (int, string, bool) {
	return 0, "todo", false
}

func (f *fwsInterface) GetDistFolder() http.FileSystem {
	return f.f.dist
}

func (f *fwsInterface) GetIndex() []byte {
	return f.f.index
}
