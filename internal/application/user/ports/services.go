package ports

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
)

type Svc interface {
	CreateEntity(username string) user.Entity
}
