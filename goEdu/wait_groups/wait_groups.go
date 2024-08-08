package main

import (
	"fmt"
	"sync"
)

func counter(idGoRoutine int, wg *sync.WaitGroup) {
	fmt.Println("ID GOROUTINE:", idGoRoutine)
	for index := range 5 {
		fmt.Printf("%d ", index+1)
	}
	fmt.Println()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for index := range 5 {
		wg.Add(1)
		go counter(index, &wg)
	}

	// wait, while will call wg.Done() for each of goroutine
	wg.Wait()
}
