package errors

import "errors"

var (
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrBossUsername  = errors.New("username cannot be b0ss")

	ErrUserAlreadyExists = errors.New("user already exists")
)
