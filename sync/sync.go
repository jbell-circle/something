package sync

import (
	"sync"

	"github.com/jbell-circle/something"
)

func NewStore() something.Store {
	return &store{sync.Map{}}
}

// store implements something.Getter using an sync map safe for concurrency.
type store struct {
	m sync.Map
}

// GetSomething implements something.Getter.
func (s *store) GetSomething(id string) (any, error) {
	if v, ok := s.m.Load(id); ok {
		return v, nil
	}
	return nil, something.ErrNotFound
}

// SetSomething implements something.Setter.
func (s *store) SetSomething(id string, v any) {
	s.m.Store(id, v)
}
