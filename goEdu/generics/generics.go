package main

import "fmt"

type oneField[T any] struct {
	field T
}

// %v is universal identifier for any type of the data
func mapPrint[K comparable, V any](Map *map[K]V) {
	fmt.Println("MAP PRINT:")
	for key, value := range *Map {
		fmt.Printf("key: %v | value: %v\n", key, value)
	}
	fmt.Println()
}

func main() {
	Map1 := map[int]string{1: "a", 2: "b", 3: "c"}
	mapPrint(&Map1)

	Map2 := map[rune]oneField[string]{'a': oneField[string]{"a"}, 'b': oneField[string]{"b"}}
	mapPrint(&Map2)

	var ptr = &Map2
	var ptrToPtr = &ptr

	fmt.Printf("meaning: %v\nptr: %p\n", Map2, ptr)
	fmt.Println("ptrToPtr:", ptrToPtr)
	//fmt.Println(*ptrToPtr, **ptrToPtr)
}
