package balance

import (
	"context"
	"errors"
	"fmt"
	accountErrors "github.com/D1sordxr/fin-eventor-lite/internal/application/account/errors"
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"github.com/D1sordxr/fin-eventor-lite/pkg/postgres/codes"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DAO struct {
	e ports.Executor
}

func NewDAO(e ports.Executor) *DAO {
	return &DAO{
		e: e,
	}
}

func (d *DAO) Get(ctx context.Context, id string) (float64, error) {
	const op = "account.balance.DAO.Get"

	query := `SELECT balance FROM accounts WHERE id = $1`

	var balance float64
	err := d.e.QueryRow(ctx, query, id).Scan(&balance)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return 0, accountErrors.ErrAccountDoesNotExist
		case errors.As(err, &pgErr) && pgErr.Code == codes.NotNullViolation:
			return 0, fmt.Errorf("%s: not null violation: %w", op, err)
		default:
			return 0, fmt.Errorf("%s: unexpected database error: %w", op, err)
		}
	}
	return balance, nil
}

func (d *DAO) Update(ctx context.Context, id string, balance float64) error {
	const op = "account.balance.DAO.Update"

	query := `UPDATE accounts SET balance = $1 WHERE id = $2`
	res, err := d.e.Exec(ctx, query, balance, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("%s: postgres error: %s (%s): %w", op, pgErr.Message, pgErr.Code, err)
		}
		return fmt.Errorf("%s: database error: %w", op, err)
	}

	if res.RowsAffected() == 0 {
		return accountErrors.ErrAccountDoesNotExist
	}

	return nil
}
