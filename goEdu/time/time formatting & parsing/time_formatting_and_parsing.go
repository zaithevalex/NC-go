package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Format(time.RFC3339)
	fmt.Println("now:", now)

	timeParse, ok := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if ok != nil {
		panic(ok.Error())
	}
	fmt.Println("parsing:", timeParse)
}
