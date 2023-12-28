package store

type Store interface {
	GetID(url string) (int, error)
	GetURL(id int) (string, error)
}
