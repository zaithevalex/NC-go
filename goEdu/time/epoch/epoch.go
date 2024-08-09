package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("seconds with starting of unix generation:", now.Unix())
	fmt.Println("nanoseconds with starting of unix generation:", now.UnixNano())
}
