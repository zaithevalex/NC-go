package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 10)
	//done := make(chan bool)

	//go func() {
	//	i := 0
	//	for {
	//		fmt.Println("\t", i)
	//		fmt.Println("<-channel:", <-channel)
	//		//index, ok := <-channel
	//		//fmt.Printf("index: %d, ok: %v\n", index, ok)
	//		//if ok {
	//		//	fmt.Println("received with index:", index)
	//		//}
	//		//if !ok {
	//		//	fmt.Println("END")
	//		//}
	//		i++
	//	}
	//}()

	for index := range 3 {
		channel <- index
	}

	//for {
	//	data, ok := <-channel
	//	if ok {
	//		fmt.Println(data, ok)
	//	}
	//}

	go func() {
		for {
			data, ok := <-channel
			if ok {
				fmt.Printf("data: %d | ok: %v\n", data, ok)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	for index := range 3 {
		channel <- index
	}

	time.Sleep(10 * time.Second)
	// after closing a channel u can't record the data
	//close(channel)
	//<-done
}
