package balance

import (
	"context"

	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/ports"
)

type UseCase struct {
	dao ports.BalanceDAO
}

func NewUseCase(dao ports.BalanceDAO) *UseCase {
	return &UseCase{
		dao: dao,
	}
}

func (uc *UseCase) GetAccountBalance(ctx context.Context, accountID string) (float64, error) {
	balance, err := uc.dao.Get(ctx, accountID)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (uc *UseCase) UpdateAccountBalance(ctx context.Context, accountID string, newBalance float64) error {
	if err := uc.dao.Update(ctx, accountID, newBalance); err != nil {
		return err
	}

	return nil
}
