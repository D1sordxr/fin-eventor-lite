package user

import "github.com/D1sordxr/fin-eventor-lite/internal/domain/user"

type Converter interface {
	EntityToModel(entity user.Entity) Model
}
