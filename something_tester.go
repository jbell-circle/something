package something

import (
	"errors"
	"testing"
)

// GetterTester_GetSomething_Success validates an implementation of a Getter returns the something when it exists.
func GetterTester_GetSomething_Success(t *testing.T, setupFunc func(id, expected string) Getter) {
	// Arrange
	id := "id"
	expected := "value"
	// A setupFunc passes the arguments your implementation-under-test needs to get ready for the test.
	getter := setupFunc(id, expected)

	// Act
	actual, _ := getter.GetSomething("id")

	// Assert
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

// GetterTester_GetSomething_NotFound validates an implementation of a Getter return a ErrNotFound error when the something does not exist.
func GetterTester_GetSomething_NotFound(t *testing.T, setupFunc func() Getter) {
	// Arrange
	getter := setupFunc()

	// Act
	_, err := getter.GetSomething("non-existent-id")

	// Assert
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound error but got %t", err)
	}
}

// StoreTester_SetAndGetSomething_Success validates an implementation of a Store can set and get something.
func StoreTester_SetAndGetSomething_Success(t *testing.T, setupFunc func() Store) {
	// Arrange
	id := "id"
	expected := "value"
	store := setupFunc()

	// Act
	store.SetSomething(id, expected)
	actual, _ := store.GetSomething(id)

	// Assert
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
