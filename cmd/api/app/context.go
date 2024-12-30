package app

import (
	"context"
	"net/http"

	"github.com/oky31/go-rest-skeleton-nca/internal/models"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *Application) contextSetUser(r *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *Application) contextGetUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(userContextKey).(*models.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
