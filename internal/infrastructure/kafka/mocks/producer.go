package mocks

import "context"

type Producer struct{}

func (Producer) Publish(_ context.Context, _ []byte) error { return nil }
