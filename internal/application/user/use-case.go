package user

import (
	"context"
	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
)

type UseCase struct {
	svc  Svc
	c    Converter
	repo Repository
}

func NewUseCase(
	svc Svc,
	c Converter,
	repo Repository,
) *UseCase {
	return &UseCase{
		svc:  svc,
		c:    c,
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, dto domain.DTO) (string, error) {
	entity := u.svc.CreateEntity(dto.Username)
	if err := u.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}
