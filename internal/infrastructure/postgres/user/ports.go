package user

import "github.com/D1sordxr/fin-eventor-lite/internal/domain/user"

type converter interface {
	EntityToModel(entity user.Entity) Model
}
