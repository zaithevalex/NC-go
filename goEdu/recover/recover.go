package main

import "fmt"

func compelledPanic1() {
	panic("compelled panic 1")
}

func compelledPanic2() {
	panic("compelled panic 2")
}

func holdOverPanic() {
	// recover gets information about error and holds over this error, printing message hold over error:
	rec := recover()
	if rec != nil {
		fmt.Println("Recovered from panic:", rec)
		return
	}
	fmt.Println("recover:", rec)
}

func main() {
	defer holdOverPanic()
	compelledPanic1()
	compelledPanic2()
}
