package user

import "github.com/google/uuid"

type Entity struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
