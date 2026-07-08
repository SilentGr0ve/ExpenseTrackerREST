package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userID uuid.UUID

var (
	userIDKey = userID{}
)

func Auth(jwtSecret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			rh := response.NewResponseHandler(w, log)

			authHeader := r.Header.Get("Authorization")

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			if tokenString == "" || tokenString == authHeader {
				rh.ErrorResponse(core_errors.ErrUnauthorized, "missing bearer token")
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				rh.ErrorResponse(core_errors.ErrUnauthorized, "invalid token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				rh.ErrorResponse(core_errors.ErrUnauthorized, "claims not parsed")
				return
			}

			subStr, ok := claims["sub"].(string)
			if !ok {
				rh.ErrorResponse(core_errors.ErrUnauthorized, "invalid claims")
				return
			}

			userUUID, err := uuid.Parse(subStr)
			if err != nil {
				rh.ErrorResponse(core_errors.ErrUnauthorized, "invalid user id")
				return
			}

			ctx = context.WithValue(ctx, userIDKey, userUUID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	id, ok := ctx.Value(userIDKey).(uuid.UUID)
	return id, ok
}

func ContextWithUserID(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(
		ctx,
		userIDKey,
		id,
	)
}
