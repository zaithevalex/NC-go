package main

import "fmt"

func main() {
	if num := 9; num < 0 { // here's number is local variable, that accessed in each of if-block
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	var num int64
	fmt.Scanf("%d", &num)
	if num%2 == 0 && num > 0 {
		fmt.Println(num, "is even")
	}
}
