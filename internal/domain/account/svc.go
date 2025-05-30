package account

import (
	"encoding/json"
	"github.com/google/uuid"
)

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

func (*Svc) CreateDepositEvent(
	accountID string,
	amount float64,
) (Event, error) {
	if err := uuid.Validate(accountID); err != nil {
		return Event{}, err
	}

	return Event{
		ID:        uuid.NewString(),
		AccountID: accountID,
		Amount:    amount,
		Type:      DepositType,
	}, nil
}

func (*Svc) PayloadEvent(
	event Event,
) ([]byte, error) {
	payload, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
