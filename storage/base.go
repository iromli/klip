package storage

import "errors"

var (
	// ErrMissingItem represents error about missing list item
	ErrMissingItem = errors.New("item not found")

	// ErrMissingList represents error about missing list
	ErrMissingList = errors.New("list not found")
)
