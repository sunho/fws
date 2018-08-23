package fws

import (
	"github.com/sunho/fws/server/runtime"
	"github.com/sunho/fws/server/store"
)

type Fws struct {
	store   store.Store
	builder runtime.Builder
	runner  runtime.Runner
}

func New() *Fws {
	return &Fws{}
}
