package main

import "fmt"

func CHANNEL_INFO(channel chan int) {
	fmt.Println("CHANNEL_INFO")
	fmt.Println("length:", len(channel))
	fmt.Println("capacity:", cap(channel))
}

func main() {
	queue := make(chan int, 5)
	queue <- 1
	queue <- 2

	CHANNEL_INFO(queue)
	close(queue)
	CHANNEL_INFO(queue)

	// elements pop from channel after each of iteration:
	for elem := range queue {
		fmt.Println(elem)
	}
	CHANNEL_INFO(queue)

	//fmt.Println("AGAIN:")
	//for elem := range queue {
	//	fmt.Println(elem)
	//}
}
