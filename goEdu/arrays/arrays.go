package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array [5]int
	fmt.Println("array: ", array)

	array[4] = 100
	fmt.Println("array: ", array)
	for index := range len(array) {
		fmt.Printf("array[%d] = %d\n", index, array[index])
	}

	fmt.Println("array length: ", len(array))

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// after initialization:
	a = [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// zero elements, starting with current index and before marked index in array
	a = [...]int{100, 3: 400, 500}
	fmt.Println(a)

	// zeroing of elements starting with 2nd to 5th and starting with 8th to 10th
	b := [...]int{0, 1, 6: 6, 7, 11: 11}
	//b := []int{0, 1, 6: 6, 7, 11: 11}
	fmt.Println("b:", b)

	// double cycles:
	var doubleCycleArray [2][3]int
	for i := range len(doubleCycleArray) {
		for j := range len(doubleCycleArray[0]) {
			doubleCycleArray[i][j] = rand.Intn(10)
		}
	}

	for i := range len(doubleCycleArray) {
		for j := range len(doubleCycleArray[0]) {
			fmt.Printf("%d ", doubleCycleArray[i][j])
		}
		fmt.Println()
	}

	var doubleCycleArray2 = [2][3]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Println(doubleCycleArray2)

	doubleCycleArray3 := [2][3]int64{{1, 2, 3}, {3, 4, 5}}
	fmt.Println(doubleCycleArray3)
}
