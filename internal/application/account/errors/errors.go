package errors

import "errors"

var (
	ErrInvalidUserID        = errors.New("invalid user ID")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrUserDoesNotExist     = errors.New("user does not exist")

	ErrInvalidAccountID    = errors.New("invalid account ID")
	ErrAccountDoesNotExist = errors.New("account does not exist")
)
