package middleware

import (
	"net/http"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"go.uber.org/zap"
)

func Logger(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get("X-Request-ID")
			ctx := r.Context()

			rw := response.NewResponseWriter(w)

			l := log.With(
				zap.String("request_id", id),
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path),
			)

			ctx = logger.ToContext(r.Context(), l)

			before := time.Now()

			next.ServeHTTP(rw, r.WithContext(ctx))

			l.Info(
				"request completed",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Now().Sub(before)),
			)

		})
	}
}
