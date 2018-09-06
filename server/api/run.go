package api

import (
	"net/http"

	"github.com/sunho/fws/server/model"
)

func (a *Api) postRun(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	confs, err := a.in.GetStore().ListBotConfig(b.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	envs, err := a.in.GetStore().ListBotEnv(b.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	rb := &model.RunBot{
		Bot:     b,
		Configs: confs,
		Envs:    envs,
	}
	err = a.in.GetRunManager().Create(rb)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
}
