package main

import (
	"fmt"
	"time"
)

func main() {
	var startTime = time.Now()

	// timer start to go after initialization:
	timer1 := time.NewTimer(2 * time.Second)
	timer2 := time.NewTimer(3 * time.Second)
	timer3 := time.NewTimer(5 * time.Second)

	// after calling
	fmt.Println("start time:", time.Since(startTime))

	go func() {
		<-timer3.C
		fmt.Println("timer 3 expired")
	}()

	go func() {
		timer4 := time.NewTimer(6 * time.Second)
		<-timer4.C
		fmt.Println("timer 4 expired")
	}()

	// also timer2 are expiring in this moment. When timer1'll expired, timer2'll has also 1s.
	timer1Time := <-timer1.C
	fmt.Printf("expired time: timer1: %v\n", timer1Time)

	fmt.Println("<-timer2.C:", time.Since(startTime))
	<-timer2.C
	fmt.Println("difference:", time.Since(startTime))

	//time.Sleep(3 * time.Second)

	// finish timer
	//timer3.Stop()
	//time.Sleep(2 * time.Second)

	// if we set new time for timer, timer gets rewrite new time and won't expired, while new time won't spent
	//timer3.Reset(5 * time.Second)
	time.Sleep(3 * time.Second)
}
