package inmemory

import "errors"

// Represents in-memory storage.
type Store map[string]interface{}

// New creates in memory storage.
func New() *Store {
	stor := make(Store)
	return &stor
}

// All returns all elements in store
func (stor Store) All() []interface{} {
	if stor == nil {
		return nil
	}
	els := make([]interface{}, len(stor))
	i := 0
	for _, v := range stor {
		els[i] = v
		i++
	}

	return els
}

// Exists check whether element with specific string id exists.
func (stor Store) Exists(id string) bool {
	_, ok := stor[id]
	return ok
}

// Add inserts item to the store. If ID is already exists, the
// error will be returned.
func (stor Store) Add(id string, el interface{}) error {
	if stor.Exists(id) {
		return errors.New("Item with specified id already exists.")
	}
	stor[id] = el
	return nil
}

// Remove removes items from store. If ID does not exists, the
// error will be return.
func (stor Store) Remove(id string) error {
	if !stor.Exists(id) {
		return errors.New("Element does not exists")
	}
	delete(stor, id)
	return nil
}

// Return element at by id. If element does not exists the
// error will be returned.
func (stor Store) Get(id string) (interface{}, error) {
	if !stor.Exists(id) {
		return nil, errors.New("Element does not exists")
	}
	return stor[id], nil
}
