package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			id := r.Header.Get("X-Request-ID")

			if id == "" {
				id = uuid.NewString()
			}

			r.Header.Set("X-Request-ID", id)
			rw.Header().Set("X-Request-ID", id)

			next.ServeHTTP(rw, r)
		})
	}
}
