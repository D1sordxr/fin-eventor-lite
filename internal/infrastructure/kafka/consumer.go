package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	r *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  "fin-eventor-lite-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &Consumer{
		r: r,
	}
}

func (c *Consumer) ReceivePayload(ctx context.Context) ([]byte, error) {
	msg, err := c.r.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}
