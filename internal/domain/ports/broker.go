package ports

import "context"

type Consumer interface {
	ReceivePayload(ctx context.Context) ([]byte, error)
}

type Producer interface {
	Publish(ctx context.Context, value []byte) error
}
