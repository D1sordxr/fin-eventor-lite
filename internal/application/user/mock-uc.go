package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type MockUseCase struct{}

func (*MockUseCase) Create(_ context.Context, dto DTO) (string, error) {
	switch dto.Username {
	case "b0ss":
		return "", ErrBossUsername
	case "":
		return "", ErrEmptyUsername
	default:
		if dto.Username == "error" {
			return "", errors.New("mock error")
		}
	}

	return uuid.NewString(), nil
}
