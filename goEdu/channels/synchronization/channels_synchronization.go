package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	var done = make(chan bool, 1)
	fmt.Println("START")
	go worker(done)
	//time.Sleep(time.Second)

	//<-done
	fmt.Println("END")
}
