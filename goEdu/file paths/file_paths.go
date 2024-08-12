package main

import (
	"fmt"
	"path/filepath"
)

func GET_RELATIVE_PATH(path1, path2 string) string {
	relativePath, ok := filepath.Rel(path1, path2)
	if ok != nil {
		panic(ok.Error())
	}
	return relativePath
}

func main() {
	path := filepath.Join("path1", "path2", "path3")
	fmt.Println(path)

	fileNameJSON := "config.json"
	fmt.Println("file resolution:", filepath.Ext(fileNameJSON))

	path1 := "a/b/c/d/e/f/g"
	path2 := "a/b/c/change/g"
	fmt.Println("relative path:", GET_RELATIVE_PATH(path1, path2))
	fmt.Println("relative path:", GET_RELATIVE_PATH(path2, path1))
}
