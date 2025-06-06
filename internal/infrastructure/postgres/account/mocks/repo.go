package mocks

import (
	"context"
	"sync"

	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/account"
)

type MockRepo struct {
	Entities map[string]domain.Entity
	m        *sync.Mutex
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		Entities: make(map[string]domain.Entity),
		m:        &sync.Mutex{},
	}
}

func (m *MockRepo) Save(_ context.Context, e domain.Entity) error {
	m.m.Lock()
	defer m.m.Unlock()

	m.Entities[e.ID.String()] = e

	return nil
}
