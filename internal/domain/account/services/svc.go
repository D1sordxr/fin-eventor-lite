package services

import (
	"encoding/json"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account/events/deposit"
	"github.com/google/uuid"
)

type Svc struct{}

func (*Svc) CreateEntity(
	userID string,
) (account.Entity, error) {
	uID, err := uuid.Parse(userID)
	if err != nil {
		return account.Entity{}, err
	}

	return account.Entity{
		ID:      uuid.New(),
		UserID:  uID,
		Balance: 0,
	}, nil
}

func (*Svc) CreateDepositEvent(
	accountID string,
	amount float64,
) (deposit.Event, error) {
	if err := uuid.Validate(accountID); err != nil {
		return deposit.Event{}, err
	}

	return deposit.Event{
		ID:        uuid.NewString(),
		AccountID: accountID,
		Amount:    amount,
		Type:      deposit.Deposit,
	}, nil
}

func (*Svc) PayloadEvent(
	event deposit.Event,
) ([]byte, error) {
	payload, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
