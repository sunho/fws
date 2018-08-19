package api

import (
	"net/http"

	"github.com/sunho/bot-registry/server/model"
)

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
		Key:      a.in.CreateInviteKey(req.Username),
	}
	_, err = a.in.GetStore().CreateUserInvite(n)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
}

func (a *Api) postUser(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Username  string `json:"username"`
		InviteKey string `jsom:"invite_key"`
		Password  string `jsom:"password"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}

	o, err := a.in.GetStore().GetUserInvite(req.Username)
	if err == nil {
		a.httpError(w, 404, nil)
		return
	}
	if req.InviteKey != o.Key {
		a.httpError(w, 403, nil)
		return
	}

	n := &model.User{
		Username: req.Username,
		Passhash: a.in.HashPassword(req.Password),
	}
	_, err = a.in.GetStore().CreateUser(n)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
}

func (a *Api) userMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			a.httpError(w, 401, nil)
			return
		}

		tok := c.Value
		id, username, ok := a.in.ParseToken(tok)
		if !ok {
			a.httpError(w, 401, nil)
			return
		}

		u, err := a.in.GetStore().GetUser(id)
		if err != nil {
			a.httpError(w, 500, err)
			return
		}

		if username != u.Username {
			a.httpError(w, 401, nil)
			return
		}

		next.ServeHTTP(w, withUser(r, u))
	}
	return http.HandlerFunc(fn)
}

func (a *Api) login(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Username string `json:"username"`
		Password string `jsom:"password"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}

	o, err := a.in.GetStore().GetUserByUsername(req.Username)
	if err == nil {
		a.httpError(w, 404, nil)
		return
	}
	if a.in.HashPassword(req.Password) != o.Passhash {
		a.httpError(w, 403, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    a.in.CreateToken(o.ID, o.Username),
		HttpOnly: true,
	})
	w.WriteHeader(201)
}
