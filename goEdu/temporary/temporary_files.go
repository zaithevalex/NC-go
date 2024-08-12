package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ASSERT(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempFile, _ := ioutil.TempFile("temporary", "tempfile.txt")
	defer os.Remove(tempFile.Name())
	fmt.Println("temporary file name:", tempFile.Name())
	_, _ = tempFile.Write([]byte("temporary file content"))

}
