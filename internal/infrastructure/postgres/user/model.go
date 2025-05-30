package user

import "github.com/google/uuid"

type Model struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
}
