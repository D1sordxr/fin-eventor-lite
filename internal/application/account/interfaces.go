package account

import (
	"context"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
)

type svc interface {
	CreateEntity(userID string) (account.Entity, error)
}

type repository interface {
	Save(ctx context.Context, entity account.Entity) error
}
