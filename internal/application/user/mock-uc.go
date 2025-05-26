package user

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
)

type MockUseCase struct{}

func (*MockUseCase) Create(_ context.Context, dto user.DTO) (string, error) {
	if dto.Username == "b0ss" {
		return "", ErrBossUsername
	}
	return "this-is-a-mock-id", nil
}
