package account

import "github.com/google/uuid"

type Entity struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}
