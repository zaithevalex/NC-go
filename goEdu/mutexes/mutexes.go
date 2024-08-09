package main

import (
	"fmt"
	"sync"
)

func iterator(counter *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	*counter++
	mutex.Unlock()
	wg.Done()
}

func test() {
	fmt.Println("TEST")
}

func main() {
	// mutex is resource, that allows to block part of code with general resource for definite (own) goals
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var counter = 0
	for index := 0; index < 10000; index++ {
		wg.Add(1)
		go iterator(&counter, &wg, &mutex)
		//time.Sleep(time.Millisecond)
	}

	//time.Sleep(3 * time.Second)
	//fmt.Println("counter:", counter)

	//go test()
	//time.Sleep(time.Second)

	wg.Wait()
	fmt.Println("counter:", counter)
}