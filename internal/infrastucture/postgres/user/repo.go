package user

import (
	"context"
	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
	"github.com/D1sordxr/fin-eventor-lite/pkg"
)

type Repository struct {
	c Converter
	e pkg.Executor
}

func NewRepository(
	c Converter,
	e pkg.Executor) *Repository {
	return &Repository{
		c: c,
		e: e,
	}
}

func (r *Repository) Save(ctx context.Context, entity domain.Entity) error {
	model := r.c.EntityToModel(entity)
	query := `INSERT INTO users (id, username) VALUES ($1, $2)`

	if _, err := r.e.Exec(ctx, query, model.ID, model.Username); err != nil {
		return err
	}

	return nil
}
