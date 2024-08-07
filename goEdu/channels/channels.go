package main

import (
	"fmt"
	"time"
)

func initializer1(messages chan string) {
	index := 2
	for index < 5 {
		messages <- "message #" + fmt.Sprint(index)
		index++
	}
}

func initializer2(messages *chan string) {
	index := 2
	for index < 5 {
		*messages <- "message #" + fmt.Sprint(index)
		index++
	}
}

func GET_CHANNEL_CONTENT(messages chan string) {
	for index := range len(messages) {
		fmt.Printf("message #%d: %s\n", index+1, <-messages)
	}
}

func main() {
	// chan is a type for marking of channel

	// default channel:
	//messages := make(chan string)

	// buffered channel is channel that has a initialization size:
	messages := make(chan string, 10)

	fmt.Println("START OF ANON FUNC")
	go func() {
		messages <- "message #1"

		// Channel can be to pass in a function by meaning or by address. It's similarly (the same).
		initializer1(messages)
		initializer2(&messages)
	}()
	time.Sleep(time.Second)
	fmt.Println("END OF ANON FUNC")
	fmt.Println("channel length:", len(messages))
	fmt.Println("channel capacity:", cap(messages))
	GET_CHANNEL_CONTENT(messages)
}
