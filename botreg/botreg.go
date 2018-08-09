package botreg

import (
	"github.com/sunho/bot-registry/runtime"
	"github.com/sunho/bot-registry/store"
)

type BotReg struct {
	store   store.Store
	builder runtime.Builder
	runner  runtime.Runner
}

func New() *BotReg {
	return &BotReg{}
}
