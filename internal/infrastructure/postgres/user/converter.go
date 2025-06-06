package user

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user"
)

type Converter struct{}

func (*Converter) EntityToModel(entity user.Entity) Model {
	return Model{
		ID:       entity.ID,
		Username: entity.Username,
	}
}
