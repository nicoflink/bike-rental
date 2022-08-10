package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	sessionHeader = "SessionID"
	UserIDKey     = "userID"
)

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get(sessionHeader)

		userID, err := uuid.Parse(sessionID)
		if err != nil {
			http.Error(w, "Invalid SessionID in header", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
