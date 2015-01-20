package main

import (
	"errors"

	"github.com/bitly/go-simplejson"
	"github.com/spf13/afero"
)

var (
	// ErrMissingItem represents error about missing list item
	ErrMissingItem = errors.New("item not found")

	// ErrMissingList represents error about missing list
	ErrMissingList = errors.New("list not found")
)

// JSONStorage represents a JSON-based storage
type JSONStorage struct {
	Filepath string
	Fs       afero.Fs
}

// readFromFile loads JSON from a file.
func (s *JSONStorage) readFromFile() (*simplejson.Json, error) {
	f, err := s.Fs.Open(s.Filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	j, err := simplejson.NewFromReader(f)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// writeToFile dumps JSON into a file.
func (s *JSONStorage) writeToFile(j *simplejson.Json) error {
	out, err := j.Encode()
	if err != nil {
		return err
	}

	f, err := s.Fs.Create(s.Filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(out); err != nil {
		return err
	}

	return nil
}

// Put creates or updates a list or list item.
func (s *JSONStorage) Put(list, name string, value interface{}) error {
	j, err := s.readFromFile()
	if err != nil {
		return err
	}

	// create list if not exists
	if _, ok := j.CheckGet(list); !ok {
		j.SetPath([]string{list}, make(map[string]string, 0))
	}

	// sets nested item
	if name != "" && value != "" {
		j.SetPath([]string{list, name}, value)
	}

	if err := s.writeToFile(j); err != nil {
		return err
	}
	return nil
}

// Get retrieves a list item.
func (s *JSONStorage) Get(list, name string) (string, error) {
	var result string

	j, err := s.readFromFile()
	if err != nil {
		return result, err
	}

	r, ok := j.Get(list).CheckGet(name)
	if !ok {
		return result, ErrMissingItem
	}

	result = r.MustString()
	return result, nil
}

// Map retrieves all values from a list.
func (s *JSONStorage) Map(list string) (map[string]interface{}, error) {
	var result map[string]interface{}

	j, err := s.readFromFile()
	if err != nil {
		return result, err
	}

	_, ok := j.CheckGet(list)
	if !ok {
		return result, ErrMissingList
	}

	result = j.GetPath(list).MustMap()
	return result, nil
}

// Delete removes list or list item if `name` is an empty string.
func (s *JSONStorage) Delete(list, name string) error {
	j, err := s.readFromFile()
	if err != nil {
		return err
	}

	if name != "" {
		if _, ok := j.GetPath(list).CheckGet(name); !ok {
			return ErrMissingItem
		}
		j.GetPath(list).Del(name)
	} else {
		if _, ok := j.CheckGet(list); !ok {
			return ErrMissingList
		}
		j.Del(list)
	}

	if err := s.writeToFile(j); err != nil {
		return err
	}
	return nil
}

// NewJSONStorage creates a JSON file to store all clips.
// If file is not exists, it will be created.
func NewJSONStorage(filepath string, fs afero.Fs) (*JSONStorage, error) {
	s := JSONStorage{
		Filepath: filepath,
		Fs:       fs,
	}
	// creates file if not exists
	if _, err := s.Fs.Stat(filepath); err != nil {
		f, err := s.Fs.Create(filepath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		if _, err := f.WriteString("{}"); err != nil {
			return nil, err
		}
	}

	return &s, nil
}
