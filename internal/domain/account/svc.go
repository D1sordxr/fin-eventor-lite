package account

import "github.com/google/uuid"

type Svc struct{}

func (*Svc) CreateEntity(
	userID string,
) (Entity, error) {
	uID, err := uuid.Parse(userID)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		ID:      uuid.New(),
		UserID:  uID,
		Balance: 0,
	}, nil
}
