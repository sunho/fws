package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
)

func (a *Api) listLog(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	runs, err := a.in.GetStore().ListBotRun(b.ID)
	if err != nil {
		a.httpError(w, r, 500, nil)
		return
	}
	a.jsonEncode(w, runs)
}

func (a *Api) getLog(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	num_ := chi.URLParam(r, "number")
	num, err := strconv.Atoi(num_)
	if err != nil {
		a.httpError(w, r, 404, nil)
		return
	}

	run, err := a.in.GetStore().GetBotRunLog(b.ID, num)
	if err != nil {
		a.httpError(w, r, 404, nil)
		return
	}
	fmt.Fprintln(w, string(run.Logged))
}

func (a *Api) getRunStatus(w http.ResponseWriter, r *http.Request) {
	s, err := a.in.GetRunManager().Status(getBot(r))
	if err == runtime.ErrNotExists {
		a.httpError(w, r, 404, nil)
		return
	} else if err != nil {
		a.httpError(w, r, 500, err)
		return
	}
	a.jsonEncode(w, s)
}

func (a *Api) postUpload(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	confs, err := a.in.GetStore().ListBotConfig(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	envs, err := a.in.GetStore().ListBotEnv(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	vols, err := a.in.GetStore().ListBotVolume(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}
	rb := &model.RunBot{
		Bot:     b,
		Configs: confs,
		Envs:    envs,
		Volumes: vols,
	}
	err = a.in.GetRunManager().Put(rb)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}
	w.WriteHeader(201)
}

func (a *Api) postRestart(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	err := a.in.GetRunManager().Restart(b)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}
	w.WriteHeader(201)
}
