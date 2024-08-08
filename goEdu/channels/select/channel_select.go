package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan int, 10)
	channel2 := make(chan int, 10)
	fmt.Println(channel1, channel2)

	go func() {
		time.Sleep(time.Second)
		channel1 <- 47
	}()

	go func(number int) {
		time.Sleep(2 * time.Second)
		//time.Sleep(time.Microsecond)
		//time.Sleep(time.Second)
		channel2 <- number
	}(99)

	//fmt.Println("START")
	//select {
	//case message := <-channel1:
	//	fmt.Println("channel1 message:", message)
	//case message := <-channel2:
	//	fmt.Println("channel2 message:", message)
	//}
	//fmt.Println("END")

	var timeStart time.Time = time.Now()
	fmt.Println("START")
	for index := range 3 {
		select {
		case message := <-channel1:
			fmt.Println("channel1 message:", message)
			fmt.Println("time channel1 :", time.Since(timeStart).Seconds())
		case message := <-channel2:
			fmt.Println("channel2 message:", message)
			fmt.Println("time channel2 :", time.Since(timeStart).Seconds())
		default:
			//time.Sleep(3 * time.Second)
			fmt.Println("All channels got messages.")
			fmt.Println("time default :", time.Since(timeStart).Seconds())
		}
		index++
	}
	fmt.Println("END")
	fmt.Println("TIME:", time.Since(timeStart).Seconds())
}
