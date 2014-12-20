package storage

import "errors"

var (
	ErrMissingItem = errors.New("item not found")
	ErrMissingList = errors.New("list not found")
)

type StorageManager interface {
	Put(list, name, value string) error
	Get(list, name string) (string, error)
	Map(list string) (map[string]interface{}, error)
	Delete(list, name string) error
}
