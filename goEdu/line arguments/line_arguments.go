package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println("amount of arguments:", len(args))
	for _, arg := range args {
		fmt.Println("arg:", arg)
	}
}
