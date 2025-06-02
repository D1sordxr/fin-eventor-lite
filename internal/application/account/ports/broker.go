package ports

import "context"

type MsgProducer interface {
	Publish(ctx context.Context, payload []byte) error
	// Publish(ctx context.Context, topic string, key string, value []byte) error
}
