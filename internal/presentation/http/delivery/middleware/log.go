package middleware

import (
	"fin-eventor-lite/pkg"
	"net/http"
	"time"
)

type LogMid struct {
	log pkg.Log
}

func NewLogMid(log pkg.Log) *LogMid {
	return &LogMid{
		log: log,
	}
}

func (m *LogMid) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			m.log.Info("Starting request...", "time", now)
			next.ServeHTTP(w, r)
			m.log.Info("Request finished.", "time-since-start", time.Since(now))
		},
	)
}
