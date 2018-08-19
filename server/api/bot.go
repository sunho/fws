package api

import (
	"net/http"

	"github.com/sunho/bot-registry/server/model"
)

func (a *Api) getBot(w http.ResponseWriter, r *http.Request) {
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
