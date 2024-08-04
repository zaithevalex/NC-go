package main

import "fmt"

func main() {
	// as the same, how C/C++ pointers
	variable := int64(187)
	fmt.Println("variable:", variable)
	fmt.Println("variable address:", &variable)

	// get variable address:
	pointer := &variable
	fmt.Println("pointer:", pointer)
	*pointer = 22
	fmt.Println("variable address:", pointer)
	fmt.Println("variable:", variable)

	// now variable is injecting in func1() and func2()
	fmt.Println("start:", variable)

	// firstly, variable injects by meaning
	func1(variable)
	fmt.Println("after injection in func1():", variable)

	// well, we can change meaning of variable if variable injects by address
	func2(&variable)
	fmt.Println("after injection in func2():", variable)
}

func func1(variable int64) {
	variable++
}

func func2(variable *int64) {
	(*variable)++
}
