package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
)

func (m *Middleware) Authenticator(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := m.log.With("middleware", "Authenticator")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			log.ErrorContext(
				ctx,
				"authorization header is empty",
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			log.ErrorContext(
				ctx,
				"invalid authorization header",
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := authorizationHeader[len("Bearer "):]
		tokenSecret := m.config.TokenSecret
		userData, err := auth.ParseToken(tokenString, tokenSecret)
		if err != nil {
			log.ErrorContext(
				ctx,
				"fail authentication",
				"error", err,
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		newCtx := context.WithValue(ctx, "user", userData)

		next.ServeHTTP(w, r.WithContext(newCtx))
	}

	return http.HandlerFunc(h)
}
