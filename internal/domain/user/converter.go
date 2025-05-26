package user

import "github.com/D1sordxr/fin-eventor-lite/internal/infrastucture/postgres/user"

type Converter struct{}

func (*Converter) EntityToModel(entity Entity) user.Model {
	return user.Model{
		ID:       entity.ID,
		Username: entity.Username,
	}
}

func (*Converter) EntityToDTO(entity Entity) DTO {
	return DTO{
		ID:       entity.ID.String(),
		Username: entity.Username,
	}
}
