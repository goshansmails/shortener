package mapstore

import (
	"sync"

	"github.com/goshansmails/shortener/internal/store/storeutils"
)

type Store struct {
	urlToID map[string]int
	idToURL map[int]string

	curID int

	mu sync.RWMutex
}

func New() *Store {
	return &Store{
		urlToID: make(map[string]int),
		idToURL: make(map[int]string),
		curID:   1,
	}
}

func (s *Store) GetID(url string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if id, found := s.urlToID[url]; found {
		return id, nil
	}

	id := s.curID
	s.curID++

	s.urlToID[url] = id
	s.idToURL[id] = url

	return id, nil
}

func (s *Store) GetURL(id int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, found := s.idToURL[id]
	if !found {
		return "", storeutils.GetIdNotFoundErr(id)
	}

	return url, nil
}
