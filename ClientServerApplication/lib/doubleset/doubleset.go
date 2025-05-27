package doubleset

import (
	"go/types"
	"sync"
)

type DoubleSet[T, V comparable] struct {
	mutex sync.Mutex
	set   map[T]map[V]types.Nil
}

func NewDoubleSet[T, V comparable]() *DoubleSet[T, V] {
	return &DoubleSet[T, V]{
		set: make(map[T]map[V]types.Nil),
	}
}

func (ds *DoubleSet[T, V]) Add(key1 T, key2 V) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	if ds.set == nil {
		ds.set = make(map[T]map[V]types.Nil)
	}
	if ds.set[key1] == nil {
		ds.set[key1] = make(map[V]types.Nil)
	}
	ds.set[key1][key2] = types.Nil{}
}

func (ds *DoubleSet[T, V]) Exists(key1 T, key2 V) bool {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	_, ex1 := ds.set[key1]
	_, ex2 := ds.set[key1][key2]
	if ex1 && ex2 {
		return true
	}

	return false
}

func (ds *DoubleSet[T, V]) Remove(key1 T, key2 V) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	delete(ds.set[key1], key2)
	if len(ds.set[key1]) == 1 {
		delete(ds.set, key1)
	}
}
