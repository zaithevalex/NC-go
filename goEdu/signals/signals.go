package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	fmt.Println("signals:", signals)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	//go func() {
	//	signal := <-signals
	//	fmt.Println()
	//	fmt.Println(signal)
	//}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("END")
}
