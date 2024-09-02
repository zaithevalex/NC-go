package main

import (
	"fmt"
	"os"
)

func main() {
	buf, err := os.ReadFile("templates/content.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("CONTENT:")
	fmt.Println(string(buf))
}
