package mapstore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreSuccess(t *testing.T) {
	type Pair struct {
		url string
		id  int
	}

	const n = 10
	pairs := make([]Pair, 0, n)
	for id := 1; id <= n; id++ {
		pairs = append(pairs, Pair{url: fmt.Sprintf("http://sample%d.ru", id), id: id})
	}

	store := New()
	size := 1

	for _, pair := range pairs {

		id, err := store.GetID(pair.url)
		assert.NoError(t, err)
		assert.Equal(t, pair.id, id)

		assert.Equal(t, size, len(store.urlToID))
		assert.Equal(t, size, len(store.idToURL))
		size++
	}

	// Save URLs again. Size doesn't change
	for _, pair := range pairs {

		id, err := store.GetID(pair.url)
		assert.NoError(t, err)
		assert.Equal(t, pair.id, id)

		assert.Equal(t, n, len(store.urlToID))
		assert.Equal(t, n, len(store.idToURL))
	}

	for _, pair := range pairs {

		url, err := store.GetURL(pair.id)
		assert.NoError(t, err)
		assert.Equal(t, pair.url, url)
	}
}

func TestStoreFail(t *testing.T) {
	store := New()

	for id := 1; id <= 10; id++ {
		_, err := store.GetURL(id)
		assert.Error(t, err)
	}
}
