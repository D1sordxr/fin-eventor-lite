package account

import "github.com/D1sordxr/fin-eventor-lite/internal/domain/account"

type converter interface {
	EntityToModel(entity account.Entity) Model
}
