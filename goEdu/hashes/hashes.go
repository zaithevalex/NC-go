package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	defaultKey := "default key"
	hashKey := sha256.New()
	hashKey.Write([]byte(defaultKey))

	fmt.Println("default string:", defaultKey)
	fmt.Printf("hash key: %x\n", hashKey.Sum(nil)) // hash sum of default key
}
