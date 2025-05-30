package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/shared/interfaces"
	"net/http"

	"github.com/D1sordxr/fin-eventor-lite/internal/application/user"
)

type UseCase interface {
	Create(ctx context.Context, dto user.DTO) (string, error)
}

type Handler struct {
	uc          UseCase
	chainer     interfaces.MidChainer
	middlewares []func(next http.Handler) http.Handler
}

func NewHandler(
	uc UseCase,
	chainer interfaces.MidChainer,
	middlewares ...func(next http.Handler) http.Handler,
) *Handler {
	return &Handler{
		uc:          uc,
		chainer:     chainer,
		middlewares: middlewares,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto user.DTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := h.uc.Create(r.Context(), dto)
	if err != nil {
		switch {

		// TODO: handle other specific errors

		case errors.Is(err, user.ErrBossUsername):
			http.Error(w, "Username 'b0ss' is reserved", http.StatusTeapot)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
	mux.HandleFunc("/users", h.chainer.WithMidChain(h.CreateUser, h.middlewares...))
}
