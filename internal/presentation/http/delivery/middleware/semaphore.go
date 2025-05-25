package middleware

import "net/http"

const (
	tokens = 50
)

type SemaphoreMid struct {
	tokens chan struct{}
}

func NewSemaphoreMid() *SemaphoreMid {
	return &SemaphoreMid{
		tokens: make(chan struct{}, tokens),
	}
}

func (m *SemaphoreMid) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			m.tokens <- struct{}{}
			defer func() { <-m.tokens }()

			next.ServeHTTP(w, r)
		},
	)
}
