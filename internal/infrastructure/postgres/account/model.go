package account

import "github.com/google/uuid"

type Model struct {
	ID      uuid.UUID `db:"id"`
	UserID  uuid.UUID `db:"user_id"`
	Balance float64   `db:"balance"`
}
