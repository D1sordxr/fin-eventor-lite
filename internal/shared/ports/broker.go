package ports

import "context"

type Consumer interface {
	ReceivePayload(ctx context.Context) ([]byte, error)
}
