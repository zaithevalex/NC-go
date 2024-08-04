package main

import "fmt"

func main() {
	slice := make([]int, 0)
	fmt.Println("slice:", slice, "| len:", len(slice), "| cap:", cap(slice))
	for index := range 10 {
		slice = append(slice, index)
		fmt.Println("slice:", slice, "| len:", len(slice), "| cap:", cap(slice))
		fmt.Println("sum:", sum(slice))
	}
}

func sum(numbers []int) int {
	sum := 0
	for index := range len(numbers) {
		sum += numbers[index]
	}
	return sum
}
