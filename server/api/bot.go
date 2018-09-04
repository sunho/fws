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

func (a *Api) getBot(w http.ResponseWriter, r *http.Request) {
	id_ := chi.URLParam(r, "bot")
	id, _ := strconv.Atoi(id_)

	b, err := a.in.GetStore().GetBot(id)
	if err != nil {
		a.httpError(w, 404, err)
		return
	}

	json.NewEncoder(w).Encode(b)
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
		Name:   req.Name,
		GitURL: req.GitURL,
	})
	if err != nil {
		a.httpError(w, 500, err)
		return
	}
	w.WriteHeader(201)
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
