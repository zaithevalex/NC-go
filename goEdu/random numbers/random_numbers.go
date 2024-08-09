package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("rand int:", rand.Intn(100))
	fmt.Println("rand float64:", rand.Float64())

	shift := 10
	fmt.Println("rand int with shift:", shift+rand.Intn(100-shift))

	// secret of random:
	randInt1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("real rand int:", randInt1.Intn(100))

	randInt2 := rand.New(rand.NewSource(47))
	randInt3 := rand.New(rand.NewSource(47))
	fmt.Println("equals PCG:")
	fmt.Println("randInt2:", randInt2.Intn(100))
	fmt.Println("randInt3:", randInt3.Intn(100))
}
