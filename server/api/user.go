package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sunho/fws/server/model"
)

func (a *Api) getUserInvite(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	key := r.URL.Query().Get("key")

	o, err := a.in.GetStore().GetUserInvite(username)
	if err != nil {
		a.httpError(w, 404, err)
		return
	}

	if o.Key != key {
		a.httpError(w, 403, nil)
		return
	}
	w.WriteHeader(200)
}

func (a *Api) postUserInvite(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	_, err := a.in.GetStore().GetUserInvite(username)
	if err == nil {
		a.httpError(w, 409, nil)
		return
	}
	_, err = a.in.GetStore().GetUserByUsername(username)
	if err == nil {
		a.httpError(w, 409, nil)
		return
	}

	n := &model.UserInvite{
		Username: username,
		Admin:    false,
		Key:      a.in.CreateInviteKey(username),
	}
	_, err = a.in.GetStore().CreateUserInvite(n)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	w.WriteHeader(201)
}
