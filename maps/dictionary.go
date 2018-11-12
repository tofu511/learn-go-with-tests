package maps

import "github.com/pkg/errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word were looking for")
	}

	return definition, nil
}