package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 15)
	for index := range 10 {
		channel <- index + 1
	}

	//for index := range len(channel) {
	//	fmt.Printf("index %d: %v\n", index+1, <-channel)
	//}

	//channel = make(chan int, 20)
	//fmt.Println("NEW:")
	//for index := range len(channel) {
	//	fmt.Printf("index %d: %v\n", index+1, <-channel)
	//}

	close(channel)
	//limiter := time.Tick(200 * time.Millisecond)
	//ticker := time.NewTicker(200 * time.Millisecond)
	//
	//var startTime = time.Now()
	//
	//fmt.Println("<-limiter:")
	//for index := range len(channel) {
	//	<-limiter
	//	fmt.Printf("time: %v | index: %d\n", time.Since(startTime), index)
	//}
	//fmt.Println()
	//
	//fmt.Println("time.Sleep():")
	//for index := range len(channel) {
	//	time.Sleep(200 * time.Millisecond)
	//	fmt.Printf("time: %v | index: %d\n", time.Since(startTime), index)
	//}
	//fmt.Println()
	//
	//fmt.Println("ticker:")
	//for index := range len(channel) {
	//	select {
	//	case <-ticker.C:
	//		fmt.Printf("time: %v | index: %d\n", time.Since(startTime), index)
	//	}
	//}

	limiter := make(chan time.Time, 3)
	requests := make(chan int, 5)

	var startTime = time.Now()
	for index := 0; index < 3; index++ {
		limiter <- time.Now()
	}

	for index := range cap(requests) {
		requests <- index
	}
	close(requests)

	go func() {
		// every 200ms limiter gets new packet of time
		for tick := range time.Tick(200 * time.Millisecond) {
			limiter <- tick
		}
	}()

	//time.Sleep(2 * time.Second)

	var lenRequests = len(requests)
	var index int
	for index = 0; index < lenRequests; index++ {
		fmt.Printf("index: %d | content: %d\n", index, <-requests)
		var timeMoment = <-limiter
		fmt.Printf("time since starting: %v\ntime limiter: %v\ntime.Now(): %v\n", time.Since(startTime).Milliseconds(), timeMoment, time.Now())
		fmt.Println("\n\n\n")
	}
}
