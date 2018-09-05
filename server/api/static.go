package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func (a *Api) fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := root.Open(r.URL.Path)
		if err != nil {
			a.getIndex(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	}))
}

func (a *Api) getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write(a.in.GetIndex())
}
