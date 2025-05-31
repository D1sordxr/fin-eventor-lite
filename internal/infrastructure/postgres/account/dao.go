package account

import (
	"context"
	"errors"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/shared/interfaces"
	"github.com/jackc/pgx/v5"
)

type DAO struct {
	e interfaces.Executor
}

func NewDAO(e interfaces.Executor) *DAO {
	return &DAO{
		e: e,
	}
}

func (d *DAO) GetByID(ctx context.Context, id string) (account.DTO, error) {
	query := `SELECT id, user_id, balance, created_at, updated_at FROM accounts WHERE id = $1`

	var model Model
	err := d.e.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.UserID,
		&model.Balance,
		// &model.CreatedAt,
		// &model.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return account.DTO{}, err // TODO: customize error for not found
		}
		return account.DTO{}, err
	}

	return account.DTO{
		ID:      model.ID.String(),
		UserID:  model.UserID.String(),
		Balance: model.Balance,
	}, nil
}
