package handlers

import (
	"context"
	"encoding/json"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account"
)

type useCase interface {
	ProcessDeposit(ctx context.Context, dto account.EventDTO) error
}

type Handler struct {
	uc useCase
}

func NewProcessor(uc useCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (p *Handler) Handle(ctx context.Context, payload []byte) error {
	var dto account.EventDTO
	err := json.Unmarshal(payload, &dto)
	if err != nil {
		return err
	}

	if err = p.uc.ProcessDeposit(ctx, dto); err != nil {
		return err
	}

	return nil
}
