package doubleset

import (
	"sync"
)

// doubleSet is a data structure that allows for the storage of two-dimensional sets
// of different comparable (T, V) data types.
type doubleSet[T, V comparable] struct {
	mutex sync.Mutex
	set   map[T]map[V]struct{}
}

// NewDoubleSet creates a two-dimensional doubleSet.
func NewDoubleSet[T, V comparable]() *doubleSet[T, V] {
	return &doubleSet[T, V]{
		set: make(map[T]map[V]struct{}),
	}
}

// Add adds an element {key1 T, key2 V} to doubleSet.
func (ds *doubleSet[T, V]) Add(key1 T, key2 V) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	if ds.set[key1] == nil {
		ds.set[key1] = make(map[V]struct{})
	}
	ds.set[key1][key2] = struct{}{}
}

// Exists checks whether an element {key1 T, key2 V} exists in the doubleSet.
func (ds *doubleSet[T, V]) Exists(key1 T, key2 V) bool {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	_, ex1 := ds.set[key1]
	_, ex2 := ds.set[key1][key2]
	if ex1 && ex2 {
		return true
	}

	return false
}

// Remove removes the specified element {key1 T, key2 V} from doubleSet.
// If the internal set is empty, then the external doubleSet with the key key1 is deleted.
func (ds *doubleSet[T, V]) Remove(key1 T, key2 V) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	delete(ds.set[key1], key2)
	if len(ds.set[key1]) == 0 {
		delete(ds.set, key1)
	}
}
