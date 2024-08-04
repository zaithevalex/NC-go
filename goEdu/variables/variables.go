package main

import (
	"fmt"
	"strconv"
)

func main() {

	// convert integer to string
	var a = strconv.FormatFloat(3.1412321, 'f', 4, 32)
	fmt.Println(a)

	var b, c = int32(1), int32(2)
	fmt.Println(b+c*2, 1, "hi")

	var d = float64(8)
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "sss"
	fmt.Println(f)

	g := f + "hi" + a
	fmt.Println(g)
}
