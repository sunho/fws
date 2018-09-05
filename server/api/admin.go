package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
)

func (a *Api) listBot(w http.ResponseWriter, r *http.Request) {
	bots, err := a.in.GetStore().ListBot()
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(bots)
}

func (a *Api) postBot(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Name   string `json:"name"`
		GitURL string `json:"git_url"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}

	_, err := a.in.GetStore().CreateBot(&model.Bot{
		Name:          req.Name,
		GitURL:        req.GitURL,
		WebhookSecret: a.in.CreateWebhookSecret(),
	})
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
}

func (a *Api) getBot(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	json.NewEncoder(w).Encode(b)
}

func (a *Api) putBot(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Name   string `json:"name"`
		GitURL string `json:"git_url"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}

	b := getBot(r)
	b.Name = req.Name
	b.GitURL = req.GitURL
	err := a.in.GetStore().UpdateBot(b)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(200)
}

func (a *Api) deleteBot(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	err := a.in.GetStore().DeleteBot(b)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(200)
}

func (a *Api) listUserInvite(w http.ResponseWriter, r *http.Request) {
	bots, err := a.in.GetStore().ListUserInvite()
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(bots)
}

func (a *Api) postUserInvite(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Username string `json:"username"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}

	_, err := a.in.GetStore().GetUserInvite(req.Username)
	if err == nil {
		a.httpError(w, 409, nil)
		return
	}
	_, err = a.in.GetStore().GetUserByUsername(req.Username)
	if err == nil {
		a.httpError(w, 409, nil)
		return
	}

	n := &model.UserInvite{
		Username: req.Username,
		Admin:    false,
		Key:      a.in.CreateInviteKey(req.Username),
	}
	_, err = a.in.GetStore().CreateUserInvite(n)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
}

func (a *Api) userMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		u, err := a.in.GetStore().GetUserByUsername(username)
		if err != nil {
			a.httpError(w, 404, nil)
			return
		}

		next.ServeHTTP(w, withUser(r, u))
	}
	return http.HandlerFunc(fn)
}

func (a *Api) listUser(w http.ResponseWriter, r *http.Request) {
	bots, err := a.in.GetStore().ListUser()
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(bots)
}

func (a *Api) postUserBot(w http.ResponseWriter, r *http.Request) {
	req := struct {
		ID int `json:"id"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}
	_, err := a.in.GetStore().GetBot(req.ID)
	if err != nil {
		a.httpError(w, 404, err)
		return
	}

	u := getUser(r)
	err = a.in.GetStore().CreateUserBot(u.ID, req.ID)
	if err != nil {
		a.httpError(w, 409, err)
		return
	}

	w.WriteHeader(201)
}
