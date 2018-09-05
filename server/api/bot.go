package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
)

func (a *Api) botMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id_ := chi.URLParam(r, "bot")
		id, _ := strconv.Atoi(id_)

		b, err := a.in.GetStore().GetBot(id)
		if err != nil {
			a.httpError(w, 404, err)
			return
		}

		next.ServeHTTP(w, withBot(r, b))
	}
	return http.HandlerFunc(fn)
}

func (a *Api) listUserBot(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	bots, err := a.in.GetStore().ListUserBot(u.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(bots)
}

func (a *Api) listBotConfig(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	confs, err := a.in.GetStore().ListBotConfig(b.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(confs)
}

func (a *Api) postBotConfig(w http.ResponseWriter, r *http.Request) {
	var req model.BotConfig
	if !a.jsonDecode(w, r, &req) {
		return
	}
	b := getBot(r)
	req.BotID = b.ID

	_, err := a.in.GetStore().CreateBotConfig(&req)
	if err != nil {
		a.httpError(w, 409, err)
		return
	}

	w.WriteHeader(201)
}

func (a *Api) putBotConfig(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (a *Api) deleteBotConfig(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) listBotVolume(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	vols, err := a.in.GetStore().ListBotVolume(b.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(vols)
}

func (a *Api) postBotVolume(w http.ResponseWriter, r *http.Request) {
	var req model.BotVolume
	if !a.jsonDecode(w, r, &req) {
		return
	}
	b := getBot(r)
	req.BotID = b.ID

	_, err := a.in.GetStore().CreateBotVolume(&req)
	if err != nil {
		a.httpError(w, 409, err)
		return
	}

	w.WriteHeader(201)
}

func (a *Api) putBotVolume(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (a *Api) deleteBotVolume(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) listBotEnv(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	envs, err := a.in.GetStore().ListBotEnv(b.ID)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(envs)
}

func (a *Api) postBotEnv(w http.ResponseWriter, r *http.Request) {
	var req model.BotEnv
	if !a.jsonDecode(w, r, &req) {
		return
	}
	b := getBot(r)
	req.BotID = b.ID

	_, err := a.in.GetStore().CreateBotEnv(&req)
	if err != nil {
		a.httpError(w, 409, err)
		return
	}

	w.WriteHeader(201)
}

func (a *Api) putBotEnv(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (a *Api) deleteBotEnv(w http.ResponseWriter, r *http.Request) {

}
