package http

import (
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
)

type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middlewares []middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return middleware.ChainMiddleware(
		r.Handler,
		r.Middlewares...,
	)
}
