package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case Ticker := <-ticker.C:
				fmt.Println("Tick at", Ticker)
			default:
				time.Sleep(100 * time.Millisecond)
				fmt.Println("     .")
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
