package storage

import "errors"

var (
	// ErrMissingItem represents error about missing list item
	ErrMissingItem = errors.New("item not found")

	// ErrMissingList represents error about missing list
	ErrMissingList = errors.New("list not found")
)

// Manager is an interface for a storage.
type Manager interface {
	Put(list, name, value string) error
	Get(list, name string) (string, error)
	Map(list string) (map[string]interface{}, error)
	Delete(list, name string) error
}
