package user

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user"
)

type converter interface {
	EntityToModel(entity user.Entity) Model
}
