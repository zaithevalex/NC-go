package main

import (
	"fmt"
	"time"
)

func cycle(s string) {
	for index := range 3 {
		fmt.Println(s, ":", index)
	}
}

func main() {
	cycle("direct")
	go cycle("goroutine #1")

	// goroutines need in time. If u don't set delay time, goroutines haven't time for executing.
	// So we always delay something method for executing of goroutines.
	// For example, one microsecond can be few for executing of goroutines, but one millisecond
	// can be already enough
	//time.Sleep(time.Microsecond)
	time.Sleep(time.Millisecond)

	// anonymous methods:
	variable := func(s string) string {
		go cycle(s)
		time.Sleep(time.Millisecond)
		return "go routine"
	}("goroutine #2")
	fmt.Println("variable:", variable)

	// goroutine for anonymous method. Starting with executing fmt.Println("END") and
	// anonymous method will call only after it. We must set time.Sleep for executing
	// of goroutines.
	go func(s string) {
		cycle(s)
	}("goroutine #3")
	fmt.Println("END")
	time.Sleep(time.Millisecond)
}
