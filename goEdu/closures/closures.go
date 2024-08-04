package main

import "fmt"

func main() {
	var variable1 = counter()
	fmt.Println("address:", variable1)
	fmt.Println("!!! call(): ", variable1())
	fmt.Println("!!! call(): ", variable1())

	variable2 := counter()
	fmt.Println("address:", variable2)
	fmt.Println("!!! call(): ", variable2())
	fmt.Println("!!! call(): ", variable2())
}

func counter() func() int {
	fmt.Println("!!! CALL INSIDE !!!")
	var index = 0
	return func() int {
		index++
		return index
	}
}
