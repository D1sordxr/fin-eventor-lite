package account

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/pb/services"
)

type DepositUseCase struct {
	client services.AccountServiceClient
	svc    depositSvc
}

func NewDepositUseCase() *DepositUseCase {
	return &DepositUseCase{}
}

func (uc *DepositUseCase) ProcessDeposit(ctx context.Context, dto EventDTO) error {
	resp, err := uc.client.GetBalance(ctx, &services.GetBalanceRequest{AccountID: dto.AccountID})
	if err != nil {
		return err
	}

	account, err := uc.svc.CreateUpdatedAccount(dto, float64(resp.Balance))
	if err != nil {
		return err
	}

	// TODO: client.UpdateBalance(...)

	return nil
}
