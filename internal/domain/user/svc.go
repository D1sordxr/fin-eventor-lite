package user

import "github.com/google/uuid"

type Svc struct{}

func (*Svc) CreateEntity(username string) Entity {
	if username == "" {
		return Entity{
			ID:       uuid.Nil,
			Username: "",
		}
	}

	return Entity{
		ID:       uuid.New(),
		Username: username,
	}
}
