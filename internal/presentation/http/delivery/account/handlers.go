package account

import (
	"context"
	"encoding/json"
	domain "github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/http/middleware"
	"net/http"
)

type useCase interface {
	Create(ctx context.Context, dto domain.DTO) (string, error)
	Deposit(ctx context.Context, dto domain.DTO) error
}

type Handler struct {
	uc          useCase
	chainer     middleware.Chainer
	middlewares []func(next http.Handler) http.Handler
}

func NewHandler(
	uc useCase,
	ch middleware.Chainer,
	m []func(next http.Handler) http.Handler,
) *Handler {
	return &Handler{
		uc:          uc,
		chainer:     ch,
		middlewares: m,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto domain.DTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	accountID, err := h.uc.Create(r.Context(), dto)
	if err != nil {
		switch {

		// TODO: handle other specific errors

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"id":      accountID,
		"message": "Account created successfully",
	})

}

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	var dto domain.DTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.uc.Deposit(r.Context(), dto); err != nil {
		switch {

		// TODO: handle other specific errors

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Deposit successful",
	})
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/accounts", h.chainer.WithMidChain(h.Create, h.middlewares...))
	mux.HandleFunc("/accounts/deposit", h.chainer.WithMidChain(h.Deposit, h.middlewares...))
}
