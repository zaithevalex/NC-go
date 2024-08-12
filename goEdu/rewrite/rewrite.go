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
	ptrReadFile, ok := os.Open("content.txt")
	ASSERT(ok)
	buffer := make([]byte, 1024)
	lengthRead, ok := ptrReadFile.Read(buffer)
	ASSERT(ok)
	fmt.Println("lengthRead:", lengthRead)
	ptrReadFile.Close()

	ptrWriteFile, ok := os.Create("rewrite.txt")
	ASSERT(ok)
	defer ptrWriteFile.Close()
	lengthWrote, ok := ptrWriteFile.Write(buffer)
	ASSERT(ok)
	fmt.Println("lengthWrote:", lengthWrote)
}
