package fws

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/sunho/fws/server/api"
	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
	"github.com/sunho/fws/server/store"
)

type Fws struct {
	stor         store.Store
	buildManager *runtime.BuildManager
	builder      runtime.Builder
	runner       runtime.Runner
	runManager   *runtime.RunManager

	config Config
	dist   http.FileSystem
	index  []byte

	server *http.Server
}

func New(stor store.Store, builder runtime.Builder,
	runner runtime.Runner, config Config) (*Fws, error) {
	f := &Fws{
		stor:       stor,
		runManager: runtime.NewRunManager(runner),
		builder:    builder,
		runner:     runner,
		config:     config,
	}
	f.buildManager = runtime.NewBuildManager(stor, builder, f.runManager)

	err := f.initDist()
	if err != nil {
		return nil, err
	}

	f.initApiServer()

	return f, nil
}

func (f *Fws) initApiServer() {
	a := api.New(&fwsInterface{f})
	handler := a.Http(f.config.Dev)
	f.server = &http.Server{
		Addr:    f.config.Addr,
		Handler: handler,
	}
}

func (f *Fws) initDist() error {
	if !f.config.Dev {
		f.dist = http.Dir(f.config.Dist)

		i, err := f.dist.Open("index.html")
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(i)
		if err != nil {
			return err
		}

		f.index = buf
	}
	return nil
}

func (f *Fws) Start() {
	go func() {
		if err := f.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	f.buildManager.Start()

	_, err := f.stor.GetUserByUsername("admin")
	if err == store.ErrNotExists {
		_, err = f.stor.CreateUserInvite(&model.UserInvite{
			Username: "admin",
			Admin:    true,
			Key:      "admin",
		})
		if err != nil {
			glog.Errorf("Creating an admin invite failed, error: %v", err)
		}
	}
}

func (f *Fws) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := f.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
