package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	idForURL map[string]int
	urlForId map[int]string

	curId int

	guard sync.RWMutex
}

func New() *Storage {
	return &Storage{
		idForURL: make(map[string]int),
		urlForId: make(map[int]string),
		curId:    1,
	}
}

func (s *Storage) GetId(url string) (int, error) {

	s.guard.Lock()
	defer s.guard.Unlock()

	if id, found := s.idForURL[url]; found {
		return id, nil
	}

	id := s.curId
	s.curId++

	s.idForURL[url] = id
	s.urlForId[id] = url

	return id, nil
}

func (s *Storage) GetURL(id int) (string, error) {

	s.guard.RLock()
	defer s.guard.RUnlock()

	url, found := s.urlForId[id]
	if !found {
		return "", errors.New("not found")
	}

	return url, nil
}
