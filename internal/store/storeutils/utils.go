package storeutils

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found URL for id")

func GetIdNotFoundErr(id int) error {
	return fmt.Errorf("%w; id = %d", ErrNotFound, id)
}
