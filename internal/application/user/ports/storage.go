package ports

import (
	"context"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user"
)

type Repository interface {
	Save(ctx context.Context, entity user.Entity) error
}
