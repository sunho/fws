package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/golang/glog"
	"github.com/sunho/fws/server/runtime"
	"github.com/sunho/fws/server/store"
)

type ApiInterface interface {
	GetStore() store.Store
	GetBuildManager() *runtime.BuildManager
	CreateInviteKey(username string) string
	ComparePassword(password string, hash string) bool
	HashPassword(password string) string
	CreateToken(id int, username string) string
	ParseToken(tok string) (int, string, bool)
	GetDistFolder() http.FileSystem
	GetIndex() []byte
}

type Api struct {
	in ApiInterface
}

func New(in ApiInterface) *Api {
	return &Api{in}
}

func (a *Api) Http() http.Handler {
	r := chi.NewRouter()
	a.cors(r)
	a.fileServer(r, "/", a.in.GetDistFolder())
	r.Route("/api", a.apiRoute)

	return r
}

func (a *Api) apiRoute(r chi.Router) {
	r.Route("/invite", func(s chi.Router) {
		s.Post("/{username}", a.postUserInvite)
		s.Get("/{username}", a.getUserInvite)
	})

	r.Route("/bot", func(s chi.Router) {
		s.Post("/", a.postBot)
		s.Put("/", a.putBot)
		s.Route("/{bot}", func(ss chi.Router) {
			ss.Use(a.botMiddleWare)
			ss.Get("/", a.getBot)
			ss.Delete("/", a.deleteBot)
			ss.Route("/build", func(sss chi.Router) {
				sss.Post("/", a.postBuild)
			})
		})
	})

	r.Post("/register", a.register)
	r.Post("/login", a.login)
}

func (a *Api) cors(r chi.Router) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
}

func (a *Api) httpError(w http.ResponseWriter, code int, org error) {
	glog.Infof("Error in http handler, code: %v org: %v", code, org)
	http.Error(w, http.StatusText(code), code)
}

func (a *Api) httpErrorWithMsg(w http.ResponseWriter, code int, msg string, org error) {
	glog.Infof("Error in http handler, code: %v msg: %s org: %v", code, msg, org)
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
