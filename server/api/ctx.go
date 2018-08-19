package api

import (
	"context"
	"net/http"

	"github.com/sunho/bot-registry/server/model"
)

var (
	userCtxKey = contextKey{"user"}
	botCtxKey  = contextKey{"bot"}
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "sunho/bot-registry/api context value " + k.name
}

func getUser(r *http.Request) *model.User {
	entry, _ := r.Context().Value(userCtxKey).(*model.User)
	return entry
}

func withUser(r *http.Request, u *model.User) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), userCtxKey, u))
	return r
}

func getBot(r *http.Request) *model.Bot {
	entry, _ := r.Context().Value(botCtxKey).(*model.Bot)
	return entry
}

func withBot(r *http.Request, b *model.Bot) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), botCtxKey, b))
	return r
}
