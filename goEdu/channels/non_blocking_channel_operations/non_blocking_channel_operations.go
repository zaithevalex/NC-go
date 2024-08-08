package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	index := 0

	var timeStart = time.Now()
	for {
		select {
		case <-tick:
			fmt.Printf("tick. | index: %d | time: %v\n", index, time.Since(timeStart))
			continue
		case <-boom:
			fmt.Printf("BOOM! | index: %d | time: %v\n", index, time.Since(timeStart))
			return
		default:
			fmt.Printf("    . | index: %d | time: %v\n", index, time.Since(timeStart))
			time.Sleep(50 * time.Millisecond)
		}
		index++
	}
}
