package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	userErrors "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user/errors"

	"github.com/D1sordxr/fin-eventor-lite/internal/application/user/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/http/middleware"
)

type UseCase interface {
	Create(ctx context.Context, dto dto.DTO) (string, error)
}

type Handler struct {
	uc          UseCase
	chainer     middleware.Chainer
	middlewares []func(next http.Handler) http.Handler
}

func NewHandler(
	uc UseCase,
	chainer middleware.Chainer,
	middlewares ...func(next http.Handler) http.Handler,
) *Handler {
	return &Handler{
		uc:          uc,
		chainer:     chainer,
		middlewares: middlewares,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.DTO
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := h.uc.Create(r.Context(), data)
	if err != nil {
		switch {
		case errors.Is(err, userErrors.ErrEmptyUsername):
			http.Error(w, "Username cannot be empty", http.StatusUnprocessableEntity)
		case errors.Is(err, userErrors.ErrBossUsername):
			http.Error(w, "Username 'b0ss' is reserved", http.StatusForbidden)
		case errors.Is(err, userErrors.ErrUserAlreadyExists):
			http.Error(w, "User already exists", http.StatusConflict)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"id":      userID,
		"message": "User created successfully",
	})
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", h.chainer.WithMidChain(h.Create, h.middlewares...))
}
