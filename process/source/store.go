package source

import (
	"errors"
	"fmt"
	"sync"
)

var store sync.Map

var Prefix = "$source"

var ErrNotFound = errors.New("cannot find the source")

var DefaultStore = NewStore()

func NewStore() *Store {
	return &Store{}
}

type Store struct {
}

// 存储资源对象
func (*Store) Store(input string) {

	sr := NewSource(input)

	alias := fmt.Sprintf("%s.%s", Prefix, sr.Name())

	store.Store(alias, sr)
}

// 查找资源对象
func (*Store) Find(alias string) (*Source, error) {

	s, ok := store.Load(alias)
	if ok {
		return s.(*Source), nil
	}
	return nil, ErrNotFound
}
