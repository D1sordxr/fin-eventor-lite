package middleware

import "net/http"

type Chainer interface {
	WithMidChain(
		handler http.HandlerFunc,
		middlewares ...func(next http.Handler) http.Handler,
	) http.HandlerFunc
}

type ChainerImpl struct{}

func (*ChainerImpl) WithMidChain(
	handler http.HandlerFunc,
	middlewares ...func(http.Handler) http.Handler,
) http.HandlerFunc {
	var h http.Handler = handler

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
