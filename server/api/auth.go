package api

import (
	"net/http"

	"github.com/sunho/fws/server/model"
)

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

func (a *Api) adminMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		u := getUser(r)
		if !u.Admin {
			a.httpError(w, 401, nil)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (a *Api) register(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Key      string `jsom:"key"`
		Username string `json:"username"`
		Nickname string `jsom:"nickname"`
		Password string `jsom:"password"`
	}{}
	if !a.jsonDecode(w, r, &req) {
		return
	}

	o, err := a.in.GetStore().GetUserInvite(req.Username)
	if err != nil {
		a.httpError(w, 404, nil)
		return
	}
	if req.Key != o.Key {
		a.httpError(w, 403, nil)
		return
	}

	_, err = a.in.GetStore().GetUserByNickname(req.Nickname)
	if err == nil {
		a.httpError(w, 409, nil)
		return
	}

	n := &model.User{
		Username: req.Username,
		Admin:    o.Admin,
		Nickname: req.Nickname,
		Passhash: a.in.HashPassword(req.Password),
	}
	_, err = a.in.GetStore().CreateUser(n)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}

	err = a.in.GetStore().DeleteUserInvite(o)
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
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
	if err != nil {
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
