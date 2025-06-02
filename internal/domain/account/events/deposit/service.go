package deposit

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
	"github.com/google/uuid"
)

type Svc struct{}

func (*Svc) CreateUpdatedAccount(event dto.EventDTO, oldBalance float64) (account.Entity, error) {
	accountID, err := uuid.Parse(event.AccountID)
	if err != nil {
		return account.Entity{}, err
	}

	return account.Entity{
		ID:      accountID,
		Balance: oldBalance + event.Amount,
	}, nil
}
