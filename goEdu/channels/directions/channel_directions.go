package main

import "fmt"

// <-chan - channel for getting of data
// chan<- - channel for sending of data

func Send(channel1 chan<- string, message string) {
	channel1 <- message
}

func Receive(channel1 <-chan string, channel2 chan<- string) {
	//message := <-channel1
	//channel2 <- message
	channel2 <- <-channel1
}

func CHANNEL_INFO(CHANNEL chan string) {
	fmt.Println("CHANNEL_INFO")
	fmt.Println("length:", len(CHANNEL))
	fmt.Println("capacity:", cap(CHANNEL))
	fmt.Println()
}

func main() {
	channel1 := make(chan string, 10)
	channel2 := make(chan string, 10)

	CHANNEL_INFO(channel1)
	CHANNEL_INFO(channel2)
	Send(channel1, "#1 message")

	CHANNEL_INFO(channel1)
	CHANNEL_INFO(channel2)
	Receive(channel1, channel2)

	CHANNEL_INFO(channel1)
	CHANNEL_INFO(channel2)
}
