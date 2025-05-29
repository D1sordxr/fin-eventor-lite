package account

import (
	"context"

	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
)

type UseCase struct {
	svc  Svc
	repo Repository
}

func NewUseCase(
	svc Svc,
) *UseCase {
	return &UseCase{
		svc: svc,
	}
}

func (uc *UseCase) Create(ctx context.Context, dto domain.DTO) (string, error) {
	entity, err := uc.svc.CreateEntity(dto.UserID)
	if err != nil {
		return "", err
	}

	if err = uc.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}
