package events

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/ports"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/pb/services"
)

type DepositUseCase struct {
	client services.AccountServiceClient
	svc    ports.DepositSvc
}

func NewDepositUseCase(
	client services.AccountServiceClient,
	svc ports.DepositSvc,
) *DepositUseCase {
	return &DepositUseCase{
		client: client,
		svc:    svc,
	}
}

func (uc *DepositUseCase) ProcessDeposit(ctx context.Context, dto dto.EventDTO) error {
	resp, err := uc.client.GetBalance(ctx, &services.GetBalanceRequest{AccountID: dto.AccountID})
	if err != nil {
		return err
	}

	account, err := uc.svc.CreateUpdatedAccount(dto, float64(resp.Balance))
	if err != nil {
		return err
	}

	_, err = uc.client.UpdateBalance(ctx, &services.UpdateBalanceRequest{
		AccountID:  account.ID.String(),
		NewBalance: float32(account.Balance),
	})

	return nil
}
