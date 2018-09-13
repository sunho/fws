package api

import (
	"io"
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

func (a *Api) devServer(r chi.Router, path string, addr string) {
	path += "*"
	r.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		r.URL.Host = a.in.GetDistAddr()
		if r.URL.Scheme == "" {
			r.URL.Scheme = "http"
		}
		req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			a.httpError(w, r, 400, err)
			return
		}
		for name, value := range r.Header {
			req.Header.Set(name, value[0])
		}
		if err != nil {
			a.httpError(w, r, 500, err)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			a.httpError(w, r, 500, err)
			return
		}

		if resp.StatusCode == 404 {
			resp2, err := http.Get(addr + "/index.html")
			if err != nil {
				a.httpError(w, r, 500, nil)
				return
			}
			io.Copy(w, resp2.Body)
			return
		}

		for k, v := range resp.Header {
			w.Header().Set(k, v[0])
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}))
}

func (a *Api) getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write(a.in.GetIndex())
}
