package middleware

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"net/http"
	"time"
)

type LogMid struct {
	log ports.Log
}

func NewLogMid(log ports.Log) *LogMid {
	return &LogMid{
		log: log,
	}
}

func (m *LogMid) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			lrw := &loggingResponseWriter{ResponseWriter: w}

			m.log.Info(
				"Request started",
				"method", r.Method,
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
				"user-agent", r.UserAgent(),
			)

			next.ServeHTTP(lrw, r)

			m.log.Info(
				"Request completed",
				"status", lrw.statusCode,
				"method", r.Method,
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
				"latency", time.Since(start).String(),
				"user-agent", r.UserAgent(),
			)
		},
	)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
