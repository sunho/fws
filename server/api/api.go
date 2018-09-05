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
	CreateWebhookSecret() string
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
	r.Get("/invite/{username}", a.getUserInvite)
	r.Post("/register", a.register)
	r.Post("/login", a.login)

	r.With(a.botMiddleWare).Post("/hook/{bot}", a.postWebhook)

	r.Route("/bot", func(s chi.Router) {
		s.Use(a.authMiddleWare)
		s.Get("/", a.listUserBot)
		s.Route("/{bot}", func(s chi.Router) {
			s.Use(a.botMiddleWare)
			s.Route("/build", func(s chi.Router) {
				s.Get("/", a.listBuild)
				s.Post("/", a.postBuild)
				s.Get("/{number}", a.getBuild)
			})
			s.Route("/volume", func(s chi.Router) {
				s.Get("/", a.listBotVolume)
				s.Post("/", a.postBotVolume)
			})
			s.Route("/config", func(s chi.Router) {
				s.Get("/", a.listBotConfig)
				s.Post("/", a.postBotConfig)
			})
			s.Route("/env", func(s chi.Router) {
				s.Get("/", a.listBotEnv)
				s.Post("/", a.postBotEnv)
			})
		})
	})

	r.Route("/admin", func(s chi.Router) {
		s.Route("/invite", func(s chi.Router) {
			s.Get("/", a.listUserInvite)
			s.Post("/", a.postUserInvite)
		})
		s.Route("/user", func(s chi.Router) {
			s.Get("/", a.listUser)
			s.Route("/{username}", func(s chi.Router) {
				s.Use(a.userMiddleWare)
				s.Route("/bot", func(s chi.Router) {
					s.Get("/", a.listUserBot)
					s.Post("/", a.postUserBot)
				})
			})
		})
		s.Route("/bot", func(s chi.Router) {
			s.Get("/", a.listBot)
			s.Post("/", a.postBot)
			s.Route("/{bot}", func(s chi.Router) {
				s.Use(a.botMiddleWare)
				s.Get("/", a.getBot)
				s.Put("/", a.putBot)
				s.Delete("/", a.deleteBot)
			})
		})
	})

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
