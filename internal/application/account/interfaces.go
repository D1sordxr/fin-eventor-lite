package account

import (
	"context"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
)

type Svc interface {
	CreateEntity(userID string) (account.Entity, error)
}

type Repository interface {
	Save(ctx context.Context, entity account.Entity) error
}
