package customtoken

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type contextKey string

const (
	usernameContextKey contextKey = "username"
)

func Middleware(tokenManager *TokenManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		username, err := tokenManager.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), usernameContextKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UsernameFromContext(ctx context.Context) (string, error) {
	username, ok := ctx.Value(usernameContextKey).(string)
	if !ok {
		return "", errors.New("no username in context")
	}

	return username, nil
}
