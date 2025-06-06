package kafka

import (
	"context"

	cfg "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/kafka"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Topic  string
	Writer *kafka.Writer
}

func NewProducer(config *cfg.Config) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(config.Brokers...),
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{
		Topic:  config.Topic,
		Writer: writer,
	}
}

func (p *Producer) Publish(ctx context.Context, value []byte) error {
	return p.Writer.WriteMessages(ctx, kafka.Message{
		Topic: p.Topic,
		Key:   nil,
		Value: value,
	})
}

func (p *Producer) Shutdown(_ context.Context) error {
	return p.Writer.Close()
}
