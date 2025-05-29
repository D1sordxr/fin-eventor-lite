package account

import (
	"context"

	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
	"github.com/D1sordxr/fin-eventor-lite/pkg"
)

type Repository struct {
	e pkg.Executor
	c converter
}

func NewRepository(
	e pkg.Executor,
	c converter,
) *Repository {
	return &Repository{
		e: e,
		c: c,
	}
}

func (r *Repository) Save(ctx context.Context, entity domain.Entity) error {
	query := ` INSERT INTO accounts (
		id,
		user_id,
		balance,
		created_at,
		updated_at,
	) VALUES ($1, $2, $3, NOW(), NOW())`

	model := r.c.EntityToModel(entity)

	if _, err := r.e.Exec(ctx, query, model); err != nil {
		return err
	}

	return nil
}
