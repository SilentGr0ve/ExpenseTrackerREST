package middleware

import "net/http"

func CORS(allowedOriginsList []string) Middleware {

	allowedOrigins := make(map[string]struct{})

	for _, origin := range allowedOriginsList {
		allowedOrigins[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			origin := r.Header.Get("Origin")
			if _, ok := allowedOrigins[origin]; ok {
				rw.Header().Set("Access-Control-Allow-Origin", origin)
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			if r.Method == http.MethodOptions {
				rw.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(rw, r)
		})
	}
}
