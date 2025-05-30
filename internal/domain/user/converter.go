package user

import "github.com/D1sordxr/fin-eventor-lite/internal/application/user"

type Converter struct{}

// TODO: create infra own converter

// func (*Converter) EntityToModel(entity Entity) user.Model {
// 	return user.Model{
// 		ID:       entity.ID,
// 		Username: entity.Username,
// 	}
// }

func (*Converter) EntityToDTO(entity Entity) user.DTO {
	return user.DTO{
		ID:       entity.ID.String(),
		Username: entity.Username,
	}
}
