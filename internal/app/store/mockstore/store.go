package mockstore

import (
	"errors"
)

type Store struct {
	idForURL map[string]int
	urlForID map[int]string
	curID    int
}

func New() *Store {
	return &Store{
		idForURL: make(map[string]int),
		urlForID: make(map[int]string),
		curID:    1,
	}
}

func (s *Store) GetID(url string) (int, error) {
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
	url, found := s.urlForID[id]
	if !found {
		return "", errors.New("not found")
	}

	return url, nil
}

func (s *Store) AddPair(url string, id int) {
	_, ok1 := s.urlForID[id]
	_, ok2 := s.idForURL[url]
	if ok1 || ok2 {
		panic("URL of ID already saved")
	}

	s.urlForID[id] = url
	s.idForURL[url] = id
}
