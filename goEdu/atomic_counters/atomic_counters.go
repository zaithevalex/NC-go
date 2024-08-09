package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//func increment(value1 *uint32, value2 uint32) {
//	atomic.AddUint32(value1, 37)
//	value2 = 47
//}

func main() {
	//var value1 uint32 = 10
	//var value2 uint32 = 10
	//
	//fmt.Printf("value1: %d, value2: %d\n", value1, value2)
	//increment(&value1, value2)
	//fmt.Printf("value1: %d, value2: %d\n", value1, value2)

	// if we'll start to increase something value by using goroutines it'll flow-unsafe, so atomic-value exists
	var wg sync.WaitGroup
	var counter int32 = 0

	for index := 0; index < 50; index++ {
		wg.Add(1)
		go func() {
			//go func(counter *int32) {

			for indexGoRoutine := 0; indexGoRoutine < 1000; indexGoRoutine++ {
				//counter++
				//*counter++
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()
		//}(&counter)
	}
	wg.Wait()
	//fmt.Println("Counter:", counter)
	fmt.Println("Counter:", counter)
}
