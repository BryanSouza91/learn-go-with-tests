package maps

const (
	ErrNotFound        = DictionaryErr("no matching key in dictionary")
	ErrKeyExists       = DictionaryErr("cannot add key as it exists already")
	ErrKeyDoesNotExist = DictionaryErr("cannot update key as it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = newValue
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
