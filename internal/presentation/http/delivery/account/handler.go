package account

import (
	"context"
	"encoding/json"
	"net/http"

	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
	"github.com/D1sordxr/fin-eventor-lite/pkg"
)

type useCase interface {
	Create(ctx context.Context, dto domain.DTO) (string, error)
}

type Handler struct {
	uc          useCase
	chainer     pkg.MidChainer
	middlewares []func(next http.Handler) http.Handler
}

func NewHandler(
	uc useCase,
	ch pkg.MidChainer,
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
