package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

func (a *Api) getBuildStatus(w http.ResponseWriter, r *http.Request) {
	status, err := a.in.GetBuildManager().Status(getBot(r))
	if err == runtime.ErrNotExists {
		a.httpError(w, 404, err)
		return
	} else if err != nil {
		a.httpError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(status)
}

func (a *Api) requestBuild(w http.ResponseWriter, b *model.Bot) {
	err := a.in.GetBuildManager().Request(b)
	if err == runtime.ErrAlreadyBuilding {
		a.httpErrorWithMsg(w, 409, "already building", err)
		return
	} else if err != nil {
		a.httpError(w, 500, err)
	}
	w.WriteHeader(201)
}

func (a *Api) postBuild(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	a.requestBuild(w, b)
}

func (a *Api) listBuild(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	bs, err := a.in.GetStore().ListBotBuild(b.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(bs)
}

func (a *Api) getBuild(w http.ResponseWriter, r *http.Request) {
	number_ := chi.URLParam(r, "number")
	number, err := strconv.Atoi(number_)
	if err != nil {
		a.httpError(w, 400, err)
		return
	}

	b := getBot(r)
	bl, err := a.in.GetStore().GetBotBuildLog(b.ID, number)
	if err != nil {
		a.httpError(w, 404, err)
		return
	}
	w.Write(bl.Logged)
}
