package api

import (
	"net/http"

	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

func (a *Api) requestBuild(w http.ResponseWriter, b *model.Bot) {
	err := a.in.GetBuildManager().Request(b)
	if err == runtime.ErrAlreadyBuilding {
		a.httpErrorWithMsg(w, 409, "already building", err)
		return
	} else if err != nil {
		a.httpError(w, 503, err)
	}
	w.WriteHeader(201)
}

func (a *Api) postBuild(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	a.requestBuild(w, b)
}
