package middleware

import (
	"net/http"
)

// AdminAuthMiddleware validates the admin API key
func AdminAuthMiddleware(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			providedKey := r.Header.Get("X-Admin-Key")

			if providedKey != apiKey {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"Unauthorized"}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
