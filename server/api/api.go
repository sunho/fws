package api

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/golang/glog"
	"github.com/sunho/fws/server/runtime"
	"github.com/sunho/fws/server/store"
)

type ApiInterface interface {
	GetStore() store.Store
	GetRunManager() *runtime.RunManager
	GetBuildManager() *runtime.BuildManager
	CreateWebhookSecret() string
	CreateInviteKey(username string) string
	ComparePassword(password string, hash string) bool
	HashPassword(password string) string
	CreateToken(id int, username string) string
	ParseToken(tok string) (int, string, bool)
	GetDistAddr() string
	GetDistFolder() http.FileSystem
	GetIndex() []byte
}

type Api struct {
	in ApiInterface
}

func New(in ApiInterface) *Api {
	return &Api{in}
}

func (a *Api) Http(dev bool) http.Handler {
	r := chi.NewRouter()
	if dev {
		a.cors(r)
		a.devServer(r, "/", a.in.GetDistAddr())
	} else {
		a.fileServer(r, "/", a.in.GetDistFolder())
	}
	r.Route("/api", a.apiRoute)

	return r
}

func (a *Api) apiRoute(r chi.Router) {
	r.Post("/register", a.register)

	r.Post("/login", a.login)
	r.Delete("/login", a.logout)

	r.Get("/invite/{username}", a.getUserInvite)

	r.With(a.authMiddleWare).Get("/user", a.getUser)

	r.With(a.botMiddleWare).Post("/hook/{bot}", a.postWebhook)

	r.Route("/bot", func(s chi.Router) {
		s.Use(a.authMiddleWare)
		s.Get("/", a.listUserBot)
		s.Route("/{bot}", func(s chi.Router) {
			s.Use(a.botMiddleWare)
			s.Use(a.userBotMiddleWare)
			s.Route("/status", func(s chi.Router) {
				s.Get("/build", a.getBuildStatus)
				s.Get("/run", a.getRunStatus)
			})
			s.Route("/build", func(s chi.Router) {
				s.Get("/", a.listBuild)
				s.Get("/{number}", a.getBuild)
				s.Post("/", a.postBuild)
			})
			s.Route("/log", func(s chi.Router) {
				s.Get("/", a.listLog)
				s.Get("/{number}", a.getLog)
			})
			s.Route("/volume", func(s chi.Router) {
				s.Get("/", a.listBotVolume)
				s.Post("/", a.postBotVolume)
				s.Patch("/{name}", a.patchBotVolume)
				s.Delete("/{name}", a.deleteBotVolume)
			})
			s.Route("/config", func(s chi.Router) {
				s.Get("/", a.listBotConfig)
				s.Post("/", a.postBotConfig)
				s.Patch("/{name}", a.patchBotConfig)
				s.Delete("/{name}", a.deleteBotConfig)
			})
			s.Route("/env", func(s chi.Router) {
				s.Get("/", a.listBotEnv)
				s.Post("/", a.postBotEnv)
				s.Patch("/{name}", a.patchBotEnv)
				s.Delete("/{name}", a.deleteBotEnv)
			})
			s.Post("/regenhook", a.postRegenHook)
			s.Post("/upload", a.postUpload)
			s.Post("/restart", a.postRestart)
		})
	})

	r.Route("/admin", func(s chi.Router) {
		s.Use(a.authMiddleWare)
		s.Use(a.adminMiddleWare)
		s.Route("/invite", func(s chi.Router) {
			s.Get("/", a.listUserInvite)
			s.Post("/", a.postUserInvite)
			s.Delete("/{username}", a.deleteUserInvite)
		})
		s.Route("/user", func(s chi.Router) {
			s.Get("/", a.listUser)
			s.Route("/{username}", func(s chi.Router) {
				s.Use(a.userMiddleWare)
				s.Get("/", a.getUser)
				s.Delete("/", a.deleteUser)
				s.Route("/bot", func(s chi.Router) {
					s.Get("/", a.listUserBot)
					s.Post("/", a.postUserBot)
					s.With(a.botMiddleWare).With(a.userBotMiddleWare).Delete("/{bot}", a.deleteUserBot)
				})
			})
		})
		s.Route("/bot", func(s chi.Router) {
			s.Get("/", a.listBot)
			s.Post("/", a.postBot)
			s.Route("/{bot}", func(s chi.Router) {
				s.Use(a.botMiddleWare)
				s.Get("/", a.getBot)
				s.Patch("/", a.patchBot)
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

func (a *Api) httpError(w http.ResponseWriter, r *http.Request, code int, org error) {
	if org != nil {
		glog.Infof("Error in http handler, code: %v org: %v (%s:%s)", code, org, r.Method, r.URL.Path)
	}
	http.Error(w, http.StatusText(code), code)
}

func (a *Api) httpErrorWithMsg(w http.ResponseWriter, r *http.Request, code int, msg string, org error) {
	if org != nil {
		glog.Infof("Error in http handler, code: %v msg: %s org: %v (%s:%s)", code, msg, org, r.Method, r.URL.Path)
	}
	http.Error(w, msg, code)
}

func (a *Api) jsonEncode(w http.ResponseWriter, i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Slice && v.Len() == 0 {
		w.Write([]byte("[]"))
		return
	}

	json.NewEncoder(w).Encode(i)
}

func (a *Api) jsonDecode(w http.ResponseWriter, r *http.Request, i interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		a.httpError(w, r, 400, nil)
		return false
	}
	return true
}
