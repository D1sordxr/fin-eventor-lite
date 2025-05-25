package user

import "context"

type MockUseCase struct{}

func (*MockUseCase) Create(_ context.Context, dto CreateUserDTO) (string, error) {
	if dto.Username == "b0ss" {
		return "", ErrBossUsername
	}
	return "this-is-a-mock-id", nil
}
