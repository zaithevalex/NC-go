package main

import "fmt"

func main() {
	fmt.Println("sum:", sum(1, 2, 3))

	sum := sumOf2Variables1stVersion(1, 2)
	fmt.Println("sum:", sum)
	sum = sumOf2Variables2ndVersion(3, 4)
	fmt.Println("sum:", sum)
}

func sum(numbers ...int) int {
	sum := 0
	for index := range numbers {
		sum += numbers[index]
	}
	return sum
}

func sumOf2Variables1stVersion(a int, b int) int {
	return a + b
}

func sumOf2Variables2ndVersion(a, b int) int {
	return a + b
}
