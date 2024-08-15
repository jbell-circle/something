package something

// Getter gets something.
type Getter interface {
	// GetSomething from Getter.
	// Returns an ErrNotFound if the something does not exist.
	GetSomething(id string) (any, error)
}

// Setter sets something.
type Setter interface {
	// SetSomething on the Setter.
	SetSomething(id string, v any)
}

// Store provides an interface for setting and getting something.
type Store interface {
	Getter
	Setter
}
