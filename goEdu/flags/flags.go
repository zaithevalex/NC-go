package main

import (
	"flag"
	"fmt"
)

func main() {
	// 1st param is name of flag
	// 2nd param is default meaning of flag
	// 3rd param is short description this flag
	wordFlags := flag.String("strFlag", "default", "a string")
	numberFlags1 := flag.Int("intFlag1", 0, "an int")
	boolFlags := flag.Bool("boolFlag", false, "a bool")

	var intVar int
	flag.IntVar(&intVar, "intFlag2", 0, "an int")

	flag.Parse()

	fmt.Println("word:", *wordFlags)
	fmt.Println("number:", *numberFlags1)
	fmt.Println("bool:", *boolFlags)
	fmt.Println("intVar1:", intVar)
	fmt.Println("flag.Args():", flag.Args())
}
