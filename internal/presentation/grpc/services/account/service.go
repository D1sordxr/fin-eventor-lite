package account

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/pb/services"
)

type useCase interface {
	GetAccountBalance(ctx context.Context, accountID string) (float64, error)
}

type Service struct {
	uc useCase
	services.UnimplementedAccountServiceServer
}

func NewService(uc useCase) *Service {
	return &Service{
		uc: uc,
	}
}

func (s *Service) GetBalance(ctx context.Context, req *services.GetBalanceRequest) (*services.GetBalanceResponse, error) {
	accountID := req.GetAccountID()

	balance, err := s.uc.GetAccountBalance(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return &services.GetBalanceResponse{
		Balance: float32(balance),
	}, nil
}
