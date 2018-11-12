package maps

import "github.com/pkg/errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}