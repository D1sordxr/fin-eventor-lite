package account

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/account"
)

type svc interface {
	CreateEntity(userID string) (account.Entity, error)
	CreateDepositEvent(accountID string, amount float64) (account.Event, error)
	PayloadEvent(event account.Event) ([]byte, error)
}

type repository interface {
	Save(ctx context.Context, entity account.Entity) error
}

type eventStore interface {
	Save(ctx context.Context, event account.Event) error
}

type msgProducer interface {
	Publish(ctx context.Context, payload []byte) error
	// Publish(ctx context.Context, topic string, key string, value []byte) error
}
