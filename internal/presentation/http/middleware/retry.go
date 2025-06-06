package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

const (
	retryCount = 5
)

type RetryMid struct{}

func (*RetryMid) RetryWithBackoff(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			lastErr := fmt.Errorf("service temporarily unavailable")
			backoff := time.Second * 1

			for i := 0; i < retryCount; i++ {
				rr := httptest.NewRecorder()
				next.ServeHTTP(rr, r)

				if rr.Code < 500 {
					for k, v := range rr.Header() {
						w.Header()[k] = v
					}
					w.WriteHeader(rr.Code)
					_, _ = rr.Body.WriteTo(w)
					return
				}

				lastErr = fmt.Errorf("attempt %d failed with status code %d", i+1, rr.Code)
				time.Sleep(backoff)
				backoff *= 2
			}

			http.Error(w, lastErr.Error(), http.StatusServiceUnavailable)
		},
	)
}
