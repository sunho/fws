package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
