package user

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
)

type svc interface {
	CreateEntity(username string) user.Entity
}

type repository interface {
	Save(ctx context.Context, entity user.Entity) error
}
