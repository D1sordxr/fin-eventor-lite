package user

import (
	"context"
	"errors"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/user/dto"
	errors2 "github.com/D1sordxr/fin-eventor-lite/internal/application/user/errors"
	"github.com/google/uuid"
)

type MockUseCase struct{}

func (*MockUseCase) Create(_ context.Context, dto dto.DTO) (string, error) {
	switch dto.Username {
	case "b0ss":
		return "", errors2.ErrBossUsername
	case "":
		return "", errors2.ErrEmptyUsername
	default:
		if dto.Username == "error" {
			return "", errors.New("mock error")
		}
	}

	return uuid.NewString(), nil
}
