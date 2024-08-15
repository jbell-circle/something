package memory

import (
	"testing"

	"github.com/jbell-circle/something"
)

func TestGetter_GetSomething_Success(t *testing.T) {
	something.GetterTester_GetSomething_Success(t, func(id, expected string) something.Getter {
		store := NewStore()
		store.SetSomething(id, expected)
		return store
	})
}

func TestGetter_GetSomething_NotFound(t *testing.T) {
	something.GetterTester_GetSomething_NotFound(t, func() something.Getter {
		return NewStore()
	})
}

func TestStore_SetAndGetSomething_Success(t *testing.T) {
	something.StoreTester_SetAndGetSomething_Success(t, func() something.Store {
		return NewStore()
	})
}
