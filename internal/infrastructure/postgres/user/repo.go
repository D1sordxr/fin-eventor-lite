package user

import (
	"context"
	"errors"
	"fmt"

	userErrors "github.com/D1sordxr/fin-eventor-lite/internal/application/user/errors"
	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
	"github.com/D1sordxr/fin-eventor-lite/pkg/postgres/codes"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	c converter
	e ports.Executor
}

func NewRepository(
	e ports.Executor,
	c converter,
) *Repository {
	return &Repository{
		c: c,
		e: e,
	}
}

func (r *Repository) Save(ctx context.Context, entity domain.Entity) error {
	const op = "user.Repository.Save"

	model := r.c.EntityToModel(entity)
	query := `INSERT INTO users (
                   id, 
                   username,
                   created_at,
                   updated_at
            	) VALUES ($1, $2, NOW(), NOW())`

	if _, err := r.e.Exec(ctx, query, model.ID, model.Username); err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == codes.UniqueViolation:
			return userErrors.ErrUserAlreadyExists
		case errors.Is(err, context.Canceled):
			return fmt.Errorf("%s: operation canceled", op)
		default:
			return fmt.Errorf("%s: failed to save user: %w", op, err)
		}
	}

	return nil
}
