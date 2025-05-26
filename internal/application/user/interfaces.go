package user

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
)

type Svc interface {
	CreateEntity(username string) user.Entity
}

type Converter interface {
	EntityToDTO(entity user.Entity) user.DTO
}

type Repository interface {
	Save(ctx context.Context, entity user.Entity) error
}
