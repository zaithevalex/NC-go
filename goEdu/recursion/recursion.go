package main

import "fmt"

func main() {
	fmt.Println("7!:", fact(7))

	// by using anonymous function
	factVariable := func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}
	fmt.Println("(factVariable) 7!:", factVariable(6))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
