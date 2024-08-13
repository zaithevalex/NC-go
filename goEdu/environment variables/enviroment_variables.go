package main

import (
	"fmt"
	"os"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	ok := os.Setenv("KEY1", "1")
	ASSERT(ok)
	ok = os.Setenv("CAT", "17")
	//ASSERT(ok)
	fmt.Println("CAT:", os.Getenv("CAT"))

	for index, key := range os.Environ() {
		fmt.Printf("index: %d -> key: %s\n", index, key)
	}
}
