package account

import (
	"context"
	"errors"
	"fmt"

	accountErrors "github.com/D1sordxr/fin-eventor-lite/internal/application/account/errors"
	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
	"github.com/D1sordxr/fin-eventor-lite/pkg/postgres/codes"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	e ports.Executor
	c converter
}

func NewRepository(
	e ports.Executor,
	c converter,
) *Repository {
	return &Repository{
		e: e,
		c: c,
	}
}

func (r *Repository) Save(ctx context.Context, entity domain.Entity) error {
	const op = "account.Repository.Save"

	query := `INSERT INTO accounts (
        id,
        user_id,
        balance,
        created_at,
        updated_at
    ) VALUES ($1, $2, $3, NOW(), NOW())`

	model := r.c.EntityToModel(entity)

	_, err := r.e.Exec(ctx, query, model.ID, model.UserID, model.Balance)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr) && pgErr.Code == codes.ForeignKeyViolation:
			return accountErrors.ErrUserDoesNotExist
		case errors.As(err, &pgErr) && pgErr.Code == codes.UniqueViolation:
			return accountErrors.ErrAccountAlreadyExists
		case errors.As(err, &pgErr) && pgErr.Code == codes.NotNullViolation:
			return fmt.Errorf("%s: required field is null: %w", op, err)
		default:
			return fmt.Errorf("%s: database error: %w", op, err)
		}
	}

	return nil
}
