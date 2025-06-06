package ports

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/account/events/deposit"
)

type Svc interface {
	CreateEntity(userID string) (account.Entity, error)
	CreateDepositEvent(accountID string, amount float64) (deposit.Event, error)
	PayloadEvent(event deposit.Event) ([]byte, error)
}

type DepositSvc interface {
	CreateUpdatedAccount(event dto.EventDTO, oldBalance float64) (account.Entity, error)
}
