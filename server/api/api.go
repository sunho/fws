package api

import (
	"encoding/json"
	"net/http"

	"github.com/sunho/fws/server/store"
)

type ApiInterface interface {
	GetStore() store.Store
	CreateInviteKey(username string) string
	HashPassword(password string) string
	CreateToken(id int, username string) string
	ParseToken(tok string) (int, string, bool)
}

type Api struct {
	in ApiInterface
}

func New(in ApiInterface) *Api {
	return &Api{in}
}

func (a *Api) Http() {
	r := chi.New()
}

func (a *Api) httpError(w http.ResponseWriter, code int, org error) {
	http.Error(w, http.StatusText(code), code)
}

func (a *Api) httpErrorWithMsg(w http.ResponseWriter, code int, msg string, org error) {
	http.Error(w, msg, code)
}

func (a *Api) jsonDecode(w http.ResponseWriter, r *http.Request, i interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		a.httpError(w, 400, nil)
		return false
	}
	return true
}
