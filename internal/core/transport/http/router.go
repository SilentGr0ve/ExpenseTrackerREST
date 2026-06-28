package httpserver

import (
	"fmt"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
)

type APIVersionRouter struct {
	prefix      string
	routes      []Route
	middlewares []middleware.Middleware
}

func NewAPIVersionRouter(version int, middlewares []middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		prefix:      fmt.Sprintf("/api/v%d", version),
		middlewares: middlewares,
	}
}

func (r *APIVersionRouter) AddRoutes(routes ...Route) {
	r.routes = append(r.routes, routes...)
}

func (r *APIVersionRouter) RegisterOn(mainMux *http.ServeMux) {
	for _, route := range r.routes {
		fullPath := r.prefix + route.Path

		handler := middleware.ChainMiddleware(
			route.WithMiddleware(),
			r.middlewares...,
		)

		mainMux.Handle(route.Method+" "+fullPath, handler)
	}
}
