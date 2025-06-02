package user

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/user/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/user/ports"
)

type UseCase struct {
	svc  ports.Svc
	repo ports.Repository
}

func NewUseCase(
	svc ports.Svc,
	repo ports.Repository,
) *UseCase {
	return &UseCase{
		svc:  svc,
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, dto dto.DTO) (string, error) {
	entity, err := u.svc.CreateEntity(dto.Username)
	if err != nil {
		return "", err
	}

	if err = u.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}
