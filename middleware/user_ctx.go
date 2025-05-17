package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Hritikpandey-ops/blog-site/utils"
	"github.com/segmentio/ksuid"
)

// key type for storing user ID in context
type contextKey string

const userIDKey contextKey = "userID"

// AuthMiddleware verifies the token and adds user ID to context.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		userID, err := utils.ExtractUserIDFromToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add userID to request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extracts userID from context.
func GetUserIDFromContext(ctx context.Context) (ksuid.KSUID, bool) {
	userID, ok := ctx.Value(userIDKey).(ksuid.KSUID)
	return userID, ok
}
