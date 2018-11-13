package maps

type Dictionary map[string]string

type DictionaryErr string

var (
	ErrNotFound = DictionaryErr("could not find the word were looking for")
	ErrWordExisting = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExisting = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExisting
	default:
		return err
	}

	return nil
}
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExisting
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}