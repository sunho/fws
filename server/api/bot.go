package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
)

func (a *Api) userBotMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		u := getUser(r)
		b := getBot(r)

		exist, err := a.in.GetStore().GetUserBot(u.ID, b.ID)
		if err != nil {
			a.httpError(w, r, 500, err)
			return
		}

		if !exist {
			a.httpError(w, r, 404, nil)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (a *Api) botMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id_ := chi.URLParam(r, "bot")
		id, _ := strconv.Atoi(id_)

		b, err := a.in.GetStore().GetBot(id)
		if err != nil {
			a.httpError(w, r, 404, err)
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
		a.httpError(w, r, 500, err)
		return
	}

	a.jsonEncode(w, bots)
}

func (a *Api) listBotConfig(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	confs, err := a.in.GetStore().ListBotConfig(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	a.jsonEncode(w, confs)
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
		a.httpError(w, r, 409, err)
		return
	}

	w.WriteHeader(201)
}

func (a *Api) putBotConfig(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (a *Api) deleteBotConfig(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	b := getBot(r)
	c, err := a.in.GetStore().GetBotConfig(b.ID, name)
	if err != nil {
		a.httpError(w, r, 404, err)
		return
	}

	err = a.in.GetStore().DeleteBotConfig(c)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	w.WriteHeader(200)
}

func (a *Api) listBotVolume(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	vols, err := a.in.GetStore().ListBotVolume(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	a.jsonEncode(w, vols)
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
		a.httpError(w, r, 409, err)
		return
	}

	w.WriteHeader(201)
}

func (a *Api) deleteBotVolume(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	b := getBot(r)
	v, err := a.in.GetStore().GetBotVolume(b.ID, name)
	if err != nil {
		a.httpError(w, r, 404, err)
		return
	}

	err = a.in.GetStore().DeleteBotVolume(v)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	w.WriteHeader(200)
}

func (a *Api) listBotEnv(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	envs, err := a.in.GetStore().ListBotEnv(b.ID)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	a.jsonEncode(w, envs)
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
		a.httpError(w, r, 409, err)
		return
	}

	w.WriteHeader(201)
}

func (a *Api) putBotEnv(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (a *Api) deleteBotEnv(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	b := getBot(r)
	e, err := a.in.GetStore().GetBotEnv(b.ID, name)
	if err != nil {
		a.httpError(w, r, 404, err)
		return
	}

	err = a.in.GetStore().DeleteBotEnv(e)
	if err != nil {
		a.httpError(w, r, 500, err)
		return
	}

	w.WriteHeader(200)
}
