package user

type Converter struct{}

// TODO: create infra own converter

// func (*Converter) EntityToModel(entity Entity) user.Model {
// 	return user.Model{
// 		ID:       entity.ID,
// 		Username: entity.Username,
// 	}
// }

func (*Converter) EntityToDTO(entity Entity) DTO {
	return DTO{
		ID:       entity.ID.String(),
		Username: entity.Username,
	}
}
