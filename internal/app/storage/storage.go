package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	idForURL map[string]int
	urlForID map[int]string

	curID int

	guard sync.RWMutex
}

func New() *Storage {
	return &Storage{
		idForURL: make(map[string]int),
		urlForID: make(map[int]string),
		curID:    1,
	}
}

func (s *Storage) GetID(url string) (int, error) {

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

func (s *Storage) GetURL(id int) (string, error) {

	s.guard.RLock()
	defer s.guard.RUnlock()

	url, found := s.urlForID[id]
	if !found {
		return "", errors.New("not found")
	}

	return url, nil
}
