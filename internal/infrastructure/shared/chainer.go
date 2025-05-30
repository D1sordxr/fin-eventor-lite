package shared

import "net/http"

type Chainer struct{}

func (*Chainer) WithMidChain(
	handler http.HandlerFunc,
	mids ...func(http.Handler) http.Handler,
) http.HandlerFunc {
	var h http.Handler = handler

	for i := len(mids) - 1; i >= 0; i-- {
		h = mids[i](h)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
