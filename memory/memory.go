package memory

import (
	"github.com/jbell-circle/something"
)

func NewStore() something.Store {
	return &store{make(map[string]any)}
}

// store implements something.Store using an in-memory map.
type store struct {
	m map[string]any
}

// GetSomething implements something.Getter.
func (s *store) GetSomething(id string) (any, error) {
	if v, ok := s.m[id]; ok {
		return v, nil
	}
	return nil, something.ErrNotFound
}

// SetSomething implements something.Setter.
func (s *store) SetSomething(id string, v any) {
	s.m[id] = v
}
