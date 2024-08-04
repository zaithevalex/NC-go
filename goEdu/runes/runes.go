package main

import "fmt"

func main() {
	// GO rune <=> C/C++ char
	const s = "Hello, 世界"
	fmt.Println("length:", len(s))
	for index, Rune := range s {
		fmt.Printf("index %d | symb: %c | hex: %x\n", index, Rune, Rune)
	}
}
