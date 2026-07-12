package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/docs"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/config"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	config      config.ServerConfig
	mux         *http.ServeMux
	log         *logger.Logger
	middlewares []middleware.Middleware
}

func NewHTTPServer(config config.ServerConfig, log *logger.Logger, middlewares ...middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		config:      config,
		mux:         http.NewServeMux(),
		log:         log,
		middlewares: middlewares,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		router.RegisterOn(s.mux)
	}
}

func (s *HTTPServer) RegisterSwagger() {
	s.mux.Handle(
		"/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
		),
	)

	s.mux.HandleFunc(
		"/swagger/doc.json",
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(docs.SwaggerInfo.ReadDoc()))
		})
}

func (s *HTTPServer) RegisterWeb(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		s.mux.Handle(pattern, route.WithMiddleware())
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {

	handler := middleware.ChainMiddleware(s.mux, s.middlewares...)

	server := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: handler,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("start HTTP server", zap.String("port", s.config.Port))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ch <- err
		}

	}()

	select {
	case err := <-ch:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		s.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown error: %w", err)
		}

		s.log.Warn("HTTP server stopped")
	}
	return nil
}
