package persistence

import "errors"

var (
	// ErrMissingResource Error to be returned is case resource is missing.
	ErrMissingResource = errors.New("missing resource")
)
