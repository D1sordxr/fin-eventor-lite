package user

import (
	"context"
)

type UseCase struct {
	svc  svc
	repo repository
}

func NewUseCase(
	svc svc,
	repo repository,
) *UseCase {
	return &UseCase{
		svc:  svc,
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, dto DTO) (string, error) {
	entity := u.svc.CreateEntity(dto.Username)
	if err := u.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}
