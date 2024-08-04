package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	switch n {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("another")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Print("it's a week")
		fmt.Println("day")
	}
}
