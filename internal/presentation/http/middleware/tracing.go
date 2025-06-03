package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const (
	traceIDKey   contextKey = "trace_id"
	requestIDKey contextKey = "request_id"
)

type TracingMid struct{}

func (*TracingMid) Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := getHeaderOrDefault(r, "X-Trace-ID", uuid.NewString())
		requestID := getHeaderOrDefault(r, "X-Request-ID", uuid.NewString())

		ctx := r.Context()
		ctx = context.WithValue(ctx, traceIDKey, traceID)
		ctx = context.WithValue(ctx, requestIDKey, requestID)

		w.Header().Set("X-Trace-ID", traceID)
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getHeaderOrDefault(r *http.Request, header, def string) string {
	if v := r.Header.Get(header); v != "" {
		return v
	}
	return def
}
