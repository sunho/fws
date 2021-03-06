package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

func (a *Api) getBuildStatus(w http.ResponseWriter, r *http.Request) {
	status, err := a.in.GetBuildManager().Status(getBot(r))
	if err == runtime.ErrNotExists {
		a.httpError(w, r, 404, err)
		return
	} else if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	a.jsonEncode(w, status)
}

func (a *Api) requestBuild(w http.ResponseWriter, r *http.Request, b *model.Bot) {
	err := a.in.GetBuildManager().Request(b)
	if err == runtime.ErrAlreadyBuilding {
		a.httpErrorWithMsg(w, r, 409, "already building", err)
		return
	} else if err != nil {
		a.httpError(w, r, 500, err)
	}
	w.WriteHeader(201)
}

func (a *Api) postBuild(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	a.requestBuild(w, r, b)
}

func (a *Api) listBuild(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	bs, err := a.in.GetStore().ListBotBuild(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}
	a.jsonEncode(w, bs)
}

func (a *Api) getBuild(w http.ResponseWriter, r *http.Request) {
	number_ := chi.URLParam(r, "number")
	number, err := strconv.Atoi(number_)
	if err != nil {
		a.httpError(w, r, 400, err)
		return
	}

	b := getBot(r)
	bl, err := a.in.GetStore().GetBotBuildLog(b.ID, number)
	if err != nil {
		a.httpError(w, r, 404, err)
		return
	}
	w.Write(bl.Logged)
}
