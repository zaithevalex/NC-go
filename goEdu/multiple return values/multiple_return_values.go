package main

import "fmt"

func main() {
	sum, sub := op(1, 2)
	fmt.Printf("sum: %d | sub: %d\n", sum, sub)

	// get only 1st value
	sum, _ = op(3, 5)
	fmt.Printf("sum: %d | sub: %d\n", sum, sub)

	// get only 2nd value
	_, sub = op(11, 5)
	fmt.Printf("sum: %d | sub: %d\n", sum, sub)

	// it doesn't work, I don't know, why ?
	//fmt.Printf("sum: %d | sub: %d\n", op(2, 2))
}

func op(a, b int) (int, int) {
	return a + b, a - b
}
