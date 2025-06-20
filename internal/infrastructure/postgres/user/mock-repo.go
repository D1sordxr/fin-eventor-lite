package user

import (
	"context"
	"sync"

	domain "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user"
)

type MockRepo struct {
	Entities []domain.Entity
	m        *sync.Mutex
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		Entities: make([]domain.Entity, 0, 100),
		m:        &sync.Mutex{},
	}
}

func (m *MockRepo) Save(_ context.Context, e domain.Entity) error {
	m.m.Lock()
	defer m.m.Unlock()

	m.Entities = append(m.Entities, e)

	return nil
}
