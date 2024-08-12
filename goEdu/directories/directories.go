package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// os.Mkdir is function for creating new directory on the same level:
	//ok := os.Mkdir("directories/dir1", 0755)
	//ASSERT(ok)

	// os.RemoveAll removes directory, that is parameter of this function and included directories in this directory:
	//ok := os.RemoveAll("directories/dir1")
	//ASSERT(ok)

	// create a empty file in the same directory:
	//ioutil.WriteFile("directories/empty.txt", []byte(""), 0644)

	// or it can create using anonymous function:
	//createFile := func(fileName string) {
	//	_ = ioutil.WriteFile(fileName, []byte(""), 0644)
	//}
	//fmt.Println(createFile)
	//createFile("directories/empty_from_function.txt")

	// create hierarchy of directories:
	//_ = os.MkdirAll("directories/dir1/dir2/dir3", 0755)
	//_ = os.MkdirAll("directories/dir1/a/b/c/d", 0755)

	readDir, _ := ioutil.ReadDir("./directories")
	for _, file := range readDir {
		fmt.Printf("file name: %s | isDir: %v\n", file.Name(), file.IsDir())
	}

	// recursive bypass:
	_ = filepath.Walk("directories", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(" ", path, info.IsDir())
		return nil
	})
}
