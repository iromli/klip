package main

import (
	"log"
	"testing"

	"github.com/spf13/afero"
	. "gopkg.in/check.v1"
)

type JSONSuite struct {
	storage *JSONStorage
}

var _ = Suite(&JSONSuite{})

func Test(t *testing.T) {
	TestingT(t)
}

func (s *JSONSuite) SetUpTest(c *C) {
	storage, err := NewJSONStorage("", &afero.MemMapFs{})
	if err != nil {
		log.Fatal(err)
	}
	s.storage = storage
}

func (s *JSONSuite) TestGet(c *C) {
	f, _ := s.storage.Fs.Create(s.storage.Filepath)
	f.Close()
	f.WriteString(`{"tests":{"foo":"bar"}}`)

	val, err := s.storage.Get("tests", "foo")
	c.Assert(err, IsNil)
	c.Assert(val, Equals, "bar")
}

func (s *JSONSuite) TestGetMissingItem(c *C) {
	val, err := s.storage.Get("tests", "foo")
	c.Assert(err, Equals, ErrMissingItem)
	c.Assert(val, Equals, "")
}

func (s *JSONSuite) TestMap(c *C) {
	f, _ := s.storage.Fs.Create(s.storage.Filepath)
	f.Close()
	f.WriteString(`{"tests":{"foo":"bar","lorem":"ipsum"}}`)

	val, err := s.storage.Map("tests")
	c.Assert(err, IsNil)
	c.Assert(val, DeepEquals, map[string]interface{}{"foo": "bar", "lorem": "ipsum"})
}

func (s *JSONSuite) TestMapMissingList(c *C) {
	val, err := s.storage.Map("linked")
	c.Assert(err, Equals, ErrMissingList)
	c.Assert(val, IsNil)
}

func (s *JSONSuite) TestDeleteList(c *C) {
	f, _ := s.storage.Fs.Create(s.storage.Filepath)
	f.Close()
	f.WriteString(`{"tests":{"foo":"bar"}}`)

	err := s.storage.Delete("tests", "")
	c.Assert(err, IsNil)
}

func (s *JSONSuite) TestDeleteItem(c *C) {
	f, _ := s.storage.Fs.Create(s.storage.Filepath)
	f.Close()
	f.WriteString(`{"tests":{"foo":"bar"}}`)

	err := s.storage.Delete("tests", "foo")
	c.Assert(err, IsNil)
}

func (s *JSONSuite) TestDeleteMissingList(c *C) {
	err := s.storage.Delete("tests", "")
	c.Assert(err, Equals, ErrMissingList)
}

func (s *JSONSuite) TestDeleteMissingItem(c *C) {
	err := s.storage.Delete("tests", "foo")
	c.Assert(err, Equals, ErrMissingItem)
}

func (s *JSONSuite) TestPutList(c *C) {
	err := s.storage.Put("random", "", "")
	c.Assert(err, IsNil)
}

func (s *JSONSuite) TestPutItem(c *C) {
	err := s.storage.Put("random", "foo", "bar")
	c.Assert(err, IsNil)
}
