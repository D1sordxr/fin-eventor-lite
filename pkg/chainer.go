package pkg

import "net/http"

type MidChainer interface {
	WithMidChain(
		handler http.HandlerFunc,
		middlewares ...func(next http.Handler) http.Handler,
	) http.HandlerFunc
}
