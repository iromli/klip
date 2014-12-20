package storage

type StorageManager interface {
	Put(list, name, value string) error
	Get(list, name string) (string, error)
	List(list string) (map[string]interface{}, error)
	Delete(list, name string) error
}
