package account

import (
	"context"
)

type UseCase struct {
	svc  svc
	repo repository
}

func NewUseCase(
	svc svc,
) *UseCase {
	return &UseCase{
		svc: svc,
	}
}

func (uc *UseCase) Create(ctx context.Context, dto DTO) (string, error) {
	entity, err := uc.svc.CreateEntity(dto.UserID)
	if err != nil {
		return "", err
	}

	if err = uc.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}
