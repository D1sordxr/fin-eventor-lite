package client

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/ports"
)

type UseCase struct {
	dao ports.DAO
}

func NewUseCase(dao ports.DAO) *UseCase {
	return &UseCase{
		dao: dao,
	}
}

func (uc *UseCase) GetAccountBalance(ctx context.Context, accountID string) (float64, error) {
	data, err := uc.dao.GetByID(ctx, accountID)
	if err != nil {
		return 0, err
	}

	return data.Balance, nil
}

func (uc *UseCase) UpdateAccountBalance(ctx context.Context, accountID string, newBalance float32) error {
	data := dto.DTO{
		ID:      accountID,
		Balance: float64(newBalance),
	}

	if err := uc.dao.Update(ctx, data); err != nil {
		return err
	}

	return nil
}
