package account

import (
	"context"
	"encoding/json"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/http/middleware"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/pb/services"
	"net/http"
	"time"
)

type ClientHandler struct {
	client      services.AccountServiceClient
	chainer     middleware.Chainer
	middlewares []func(next http.Handler) http.Handler
}

func NewClientHandler(
	client services.AccountServiceClient,
	ch middleware.Chainer,
	m ...func(next http.Handler) http.Handler,
) *ClientHandler {
	return &ClientHandler{
		client:      client,
		chainer:     ch,
		middlewares: m,
	}
}

func (h *ClientHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	accountID := r.URL.Query().Get("account_id")
	resp, err := h.client.GetBalance(ctx, &services.GetBalanceRequest{AccountID: accountID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]float32{
		"balance": resp.GetBalance(),
	})
}

func (h *ClientHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/accounts/balance", h.chainer.WithMidChain(h.GetBalance, h.middlewares...))
}
