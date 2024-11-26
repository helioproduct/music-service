package middleware

import (
	"music-service/pkg/logger"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logging(logger logger.Logger) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Info("Started request", "method", r.Method, "url", r.URL.Path)

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			logger.Info("Completed request", "method", r.Method, "url", r.URL.Path, "duration", duration)
		})
	}
}
