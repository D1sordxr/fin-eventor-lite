package account

import "github.com/google/uuid"

type Model struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance float64
}
