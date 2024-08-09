package main

import (
	"fmt"
	"time"
)

func closeChannel(channel chan int) {
	if channel == nil {
		panic("channel is nil")
	}
	close(channel)
	fmt.Println("channel closed")
}

func initializeChannel(channel chan int, number int) {
	channel <- number
}

func main() {
	//defer fmt.Println("defer message")
	//
	//slice := []int{10, 20, 30, 40, 50}
	//for index, elem := range slice {
	//	fmt.Printf("index #%d: %d\n", index+1, elem)
	//}
	//panic("panic message")

	channel := make(chan int, 32)
	defer closeChannel(channel)

	for i := 0; i < 10; i++ {
		go initializeChannel(channel, i)
	}

	time.Sleep(time.Second)
	fmt.Println("length channel:", len(channel))

	//fmt.Println(strings.Join([]string{"a", "b", "c"}, "_-_"))
}
