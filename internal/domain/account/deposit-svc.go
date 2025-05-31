package account

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account"
	"github.com/google/uuid"
)

type DepositSvc struct{}

func (*DepositSvc) CreateUpdatedAccount(event account.EventDTO, oldBalance float64) (Entity, error) {
	accountID, err := uuid.Parse(event.AccountID)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		ID:      accountID,
		Balance: oldBalance + event.Amount,
	}, nil
}
