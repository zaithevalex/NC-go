package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func arrayRandomInitialize(array *[50]int) {
	fmt.Printf("array address: %p\n", array)
	fmt.Println("array:", *array)
	for index := 0; index < 50; index++ {
		array[index] = rand.Intn(100) // <=> (*array)[index] = rand.Intn(100)
	}
}

func sliceRandomInitialize(slice *[]int) {
	fmt.Printf("slice address: %p\n", slice)
	for index := 0; index < 50; index++ {
		*slice = append(*slice, rand.Intn(100))
	}
}

// bubble sort:
func sortArray(array *[50]int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func main() {
	slice := make([]int, 0)
	fmt.Println(slice)
	sliceRandomInitialize(&slice)
	fmt.Println("SLICE:")
	fmt.Println("BEFORE SORTING:", slice)
	slices.Sort(slice)
	fmt.Println("AFTER SORTING:", slice)
	fmt.Println()

	array := [50]int{}
	fmt.Println("ARRAY:")
	arrayRandomInitialize(&array)
	fmt.Println("BEFORE SORTING:", array)
	sortArray(&array)
	fmt.Println("AFTER SORTING:", array)
}
