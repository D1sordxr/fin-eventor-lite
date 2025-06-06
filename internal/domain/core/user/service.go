package user

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user/errors"
	"github.com/google/uuid"
)

type Svc struct{}

func (*Svc) CreateEntity(username string) (Entity, error) {
	switch username {
	case "":
		return Entity{}, errors.ErrEmptyUsername
	case "b0ss":
		return Entity{}, errors.ErrBossUsername
	}

	return Entity{
		ID:       uuid.New(),
		Username: username,
	}, nil
}
