package mockstore

import (
	"errors"
)

type Store struct {
	urlToID map[string]int
	idToURL map[int]string
	curID   int
}

func New() *Store {
	return &Store{
		urlToID: make(map[string]int),
		idToURL: make(map[int]string),
		curID:   1,
	}
}

func (s *Store) GetID(url string) (int, error) {
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
	url, found := s.idToURL[id]
	if !found {
		return "", errors.New("not found")
	}

	return url, nil
}

func (s *Store) AddPair(url string, id int) {
	_, ok1 := s.idToURL[id]
	_, ok2 := s.urlToID[url]
	if ok1 || ok2 {
		panic("URL of ID already saved")
	}

	s.idToURL[id] = url
	s.urlToID[url] = id
}
