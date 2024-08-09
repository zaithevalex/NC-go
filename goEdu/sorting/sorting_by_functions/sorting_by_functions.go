package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func sliceRandomInitialize(slice *[]int) {
	fmt.Printf("slice address: %p\n", slice)
	for index := 0; index < 50; index++ {
		*slice = append(*slice, rand.Intn(100))
	}
}

func main() {
	slice := make([]int, 0)
	sliceRandomInitialize(&slice)

	sortedFuncByDecrease := func(elem1, elem2 int) int {
		return elem2 - elem1
	}
	fmt.Println("BEFORE SORTING:", slice)
	// as the same, that is comparator in 2nd parameter
	slices.SortFunc(slice, sortedFuncByDecrease)
	fmt.Println("AFTER SORTING:", slice)

	// same, we can sort structs by one of the them parameter
}
