package botreg

import (
	"github.com/sunho/bot-registry/server/runtime"
	"github.com/sunho/bot-registry/server/store"
)

type BotReg struct {
	store   store.Store
	builder runtime.Builder
	runner  runtime.Runner
}

func New() *BotReg {
	return &BotReg{}
}
