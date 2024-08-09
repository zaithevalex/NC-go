package main

import (
	"fmt"
	"strconv"
)

func main() {
	float64Parse, _ := strconv.ParseFloat("11.23123", 64)
	fmt.Println("float64Parse:", float64Parse)

	// base is rank of the calculation system
	int64Parse, _ := strconv.ParseInt("AAA", 16, 64)
	fmt.Println("int64Parse:", int64Parse)

	// hex parsing:
	hexParse, _ := strconv.ParseInt("0xDEADBEE", 0, 64)
	fmt.Println("hexParse:", hexParse)

	base10AtoiParse, _ := strconv.Atoi("812") //  <=> strconv.ParseInt("812", 10, 4)
	fmt.Println("base10AtoiParse:", base10AtoiParse)

	wordAtoiParse, _ := strconv.Atoi("wat") // BROKEN...
	fmt.Println("wordAtoiParse:", wordAtoiParse)
}