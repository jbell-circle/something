package something

import "errors"

var (
	// ErrNotFound indicates the something was not found.
	ErrNotFound = errors.New("not found")
)
