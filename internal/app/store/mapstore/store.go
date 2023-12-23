package mapstore

import (
	"errors"
	"sync"
)

type Store struct {
	idForURL map[string]int
	urlForID map[int]string

	curID int

	guard sync.RWMutex
}

func New() *Store {
	return &Store{
		idForURL: make(map[string]int),
		urlForID: make(map[int]string),
		curID:    1,
	}
}

func (s *Store) GetID(url string) (int, error) {

	s.guard.Lock()
	defer s.guard.Unlock()

	if id, found := s.idForURL[url]; found {
		return id, nil
	}

	id := s.curID
	s.curID++

	s.idForURL[url] = id
	s.urlForID[id] = url

	return id, nil
}

func (s *Store) GetURL(id int) (string, error) {

	s.guard.RLock()
	defer s.guard.RUnlock()

	url, found := s.urlForID[id]
	if !found {
		return "", errors.New("not found")
	}

	return url, nil
}
