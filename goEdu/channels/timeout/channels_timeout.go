package main

import (
	"fmt"
	"time"
)

func channelInitializer(channel chan int, number int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	channel <- number
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go channelInitializer(channel1, 12, 1)
	go channelInitializer(channel2, 47, 2)

	var timeStart = time.Now()
	//select {
	//case number := <-channel1:
	//	fmt.Println("channel1 received:", number)
	//case number := <-channel2:
	//	fmt.Println("channel2 received:", number)
	//case <-time.After(750 * time.Millisecond):
	//	fmt.Println("timeout #1")
	//}
	//fmt.Println("select #1:", time.Since(timeStart))
	//fmt.Println()

	//select {
	//case number := <-channel1:
	//	fmt.Println("channel1 received:", number)
	//case number := <-channel2:
	//	fmt.Println("channel2 received:", number)
	//case <-time.After(750 * time.Millisecond):
	//	fmt.Println("timeout #2")
	//}
	//fmt.Println("select #2:", time.Since(timeStart))
	//fmt.Println()
	//
	//select {
	//case number := <-channel1:
	//	fmt.Println("channel1 received:", number)
	//case number := <-channel2:
	//	fmt.Println("channel2 received:", number)
	//case <-time.After(750 * time.Millisecond):
	//	fmt.Println("TRIGGER:", time.Since(timeStart).Seconds())
	//	fmt.Println("timeout #3")
	//}
	//fmt.Println("select #3:", time.Since(timeStart))
	//fmt.Println()
	//
	//select {
	//case number := <-channel1:
	//	fmt.Println("channel1 received:", number)
	//case number := <-channel2:
	//	fmt.Println("channel2 received:", number)
	//case <-time.After(750 * time.Millisecond):
	//	fmt.Println("timeout #4")
	//}
	//fmt.Println("select #4:", time.Since(timeStart))
	//fmt.Println()
	//
	//select {
	//case number := <-channel1:
	//	fmt.Println("channel1 received:", number)
	//case number := <-channel2:
	//	fmt.Println("channel2 received:", number)
	//case <-time.After(750 * time.Millisecond):
	//	fmt.Println("timeout #5")
	//}
	//fmt.Println("select #5:", time.Since(timeStart))
	//fmt.Println()

	for index := range 5 {
		select {
		case number := <-channel1:
			fmt.Println("channel1 received:", number)
		case number := <-channel2:
			fmt.Println("channel2 received:", number)
		case <-time.After(750 * time.Millisecond):
			fmt.Printf("timeout #%d\n", index+1)
		}
		fmt.Printf("select #%d: %v\n", index+1, time.Since(timeStart))
		fmt.Println()
		index++
	}
}
