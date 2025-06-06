package handlers

import (
	"context"
	"encoding/json"

	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
)

type useCase interface {
	ProcessDeposit(ctx context.Context, dto dto.EventDTO) error
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
	var data dto.EventDTO
	err := json.Unmarshal(payload, &data)
	if err != nil {
		return err
	}

	if err = p.uc.ProcessDeposit(ctx, data); err != nil {
		return err
	}

	return nil
}
