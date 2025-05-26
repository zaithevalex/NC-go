package lib

import "go/types"

type DoubleSet[T, V comparable] struct {
	set1 map[T]types.Nil
	set2 map[V]types.Nil
}

func NewDoubleSet[T, V comparable]() *DoubleSet[T, V] {
	return &DoubleSet[T, V]{
		set1: make(map[T]types.Nil),
		set2: make(map[V]types.Nil),
	}
}

func (ds *DoubleSet[T, V]) Add(key1 T, key2 V) {
	ds.set1[key1] = types.Nil{}
	ds.set2[key2] = types.Nil{}
}

func (ds *DoubleSet[T, V]) Exists(key1 T, key2 V) bool {
	_, ex1 := ds.set1[key1]
	_, ex2 := ds.set2[key2]
	if ex1 && ex2 {
		return true
	}

	return false
}

func (ds *DoubleSet[T, V]) Remove(key1 T, key2 V) {
	delete(ds.set1, key1)
	delete(ds.set2, key2)
}

func (ds *DoubleSet[T, V]) Size() int {
	return len(ds.set1)
}
