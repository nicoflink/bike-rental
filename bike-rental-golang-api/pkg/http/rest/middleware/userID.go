package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/nicoflink/bike-rental/pkg/http/rest/errors"
)

const (
	sessionHeader = "SessionID"
	UserIDKey     = "userID"
)

// UserCtx extracts the sessionID (=UserID) from the http header and puts it into the context.
// Each request must contain a valid sessionID in form of an uuid4.
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get(sessionHeader)

		userID, err := uuid.Parse(sessionID)
		if err != nil {
			http.Error(w, errors.ErrSessionID, http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
