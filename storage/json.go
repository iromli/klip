package storage

import (
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/bitly/go-simplejson"
)

// JSONStorage represents a JSON-based storage
type JSONStorage struct {
	Filepath string
}

func (s *JSONStorage) getJSON() (*simplejson.Json, error) {
	buf, err := ioutil.ReadFile(s.Filepath)
	if err != nil {
		return nil, err
	}

	j, err := simplejson.NewJson(buf)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// Put creates or updates a list or list item.
func (s *JSONStorage) Put(list, name, value string) error {
	j, err := s.getJSON()
	if err != nil {
		return err
	}

	// create list if not exists
	if _, ok := j.CheckGet(list); !ok {
		j.SetPath([]string{list}, make(map[string]string, 0))
	}

	if name != "" && value != "" {
		j.SetPath([]string{list, name}, value)
	}

	out, err := j.Encode()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(s.Filepath, out, 0644); err != nil {
		return err
	}
	return nil
}

// Get retrieves a list item.
func (s *JSONStorage) Get(list, name string) (string, error) {
	var result string

	j, err := s.getJSON()
	if err != nil {
		return result, err
	}

	r, ok := j.Get(list).CheckGet(name)
	if !ok {
		return result, errors.New("item not found")
	}

	result = r.MustString()
	return result, nil
}

// List retrieves all values from a list.
func (s *JSONStorage) List(list string) (map[string]interface{}, error) {
	var result map[string]interface{}

	j, err := s.getJSON()
	if err != nil {
		return result, err
	}

	_, ok := j.CheckGet(list)
	if !ok {
		return result, errors.New("list not found")
	}

	result = j.GetPath(list).MustMap()
	return result, nil
}

// Delete removes list or list item if `name` is an empty string.
func (s *JSONStorage) Delete(list, name string) error {
	return nil
}

// NewJSONStorage creates a JSON file to store all clips.
// This file is located under current user's home directory, e.g. `/home/user/.clip`.
// If file is not exists, it will be created.
func NewJSONStorage() (*JSONStorage, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	filepath := path.Join(u.HomeDir, ".clip")

	if _, err := os.Stat(filepath); err != nil {
		f, err := os.Create(filepath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		if _, err := f.WriteString("{}"); err != nil {
			return nil, err
		}
	}

	s := JSONStorage{Filepath: filepath}
	return &s, nil
}
