package middlewares

import (
	"bls/logger"
	"net/http"
	"time"

	"github.com/elisiei/zlog"
)

func LoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		logger.Log.Debugw("request", zlog.F{
			"method":     r.Method,
			"path":       r.URL.Path,
			"duration":   time.Since(start).String(),
			"remote":     r.RemoteAddr,
			"user-agent": r.UserAgent(),
		})
	})
}
