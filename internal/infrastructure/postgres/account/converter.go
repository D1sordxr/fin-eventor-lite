package account

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/core/account"
)

type Converter struct{}

func (*Converter) EntityToModel(entity account.Entity) Model {
	return Model{
		ID:      entity.ID,
		UserID:  entity.UserID,
		Balance: entity.Balance,
	}
}
