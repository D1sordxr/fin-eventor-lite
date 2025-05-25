package middleware

import "net/http"

type MethodMid struct {
	Method string
}

func NewMethodMid(method string) *MethodMid {
	return &MethodMid{Method: method}
}

func (m *MethodMid) OnlyPost(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m.Method {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}
