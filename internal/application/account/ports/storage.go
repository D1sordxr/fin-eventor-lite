package ports

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account/events/deposit"
)

type Repository interface {
	Save(ctx context.Context, entity account.Entity) error
}

type EventStore interface {
	Save(ctx context.Context, event deposit.Event) error
}

type DAO interface {
	GetByID(ctx context.Context, id string) (dto.DTO, error)
	Update(ctx context.Context, data dto.DTO) error
}
