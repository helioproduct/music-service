package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

func Panic(logger *slog.Logger) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("recover middleware", "error", err)
					logger.Debug(string(debug.Stack()))
					http.Error(w, "Internal server error", 500)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
