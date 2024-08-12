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
	buffer := "Hello, Aleksandr Aleksandr Zaitcev !"
	//ok := ioutil.WriteFile("file.txt", []byte(buffer), 0644)
	//ASSERT(ok)

	ptrFile, ok := os.Create("file.txt")
	ASSERT(ok)
	defer ptrFile.Close()

	lengthWrote, ok := ptrFile.Write([]byte(buffer))
	ASSERT(ok)
	fmt.Printf("Wrote %d bytes\n", lengthWrote)

	byteBuffer := []byte{115, 111, 109, 101, 10}
	ptrFile1, ok := os.Create("file1.txt")
	ASSERT(ok)
	defer ptrFile1.Close()
	_, _ = ptrFile1.Write(byteBuffer)

	ptrFile2, ok := os.Create("file1.txt")
	ASSERT(ok)
	defer ptrFile2.Close()
	_, _ = ptrFile2.WriteString(buffer)
	_ = ptrFile2.Sync()
}
