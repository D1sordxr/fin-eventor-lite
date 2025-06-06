package consumer

import (
	"context"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
)

type handler interface {
	Handle(ctx context.Context, payload []byte) error
}

type Consumer struct {
	c ports.Consumer
	h handler
}

func NewConsumer(c ports.Consumer, h handler) *Consumer {
	return &Consumer{
		c: c,
		h: h,
	}
}

func (c *Consumer) StartConsumer(ctx context.Context) error {
	for {
		payload, err := c.c.ReceivePayload(ctx)
		if err != nil {
			return err
		}

		if err = c.h.Handle(ctx, payload); err != nil {
			return err
		}
	}
}
