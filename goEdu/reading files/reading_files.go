package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1st way of reading file:
	//readBuf, ok := ioutil.ReadFile("file.txt")
	//if ok != nil {
	//	panic(ok.Error())
	//}
	//fmt.Println("content:", string(readBuf))

	// 2nd way of reading file:
	//readBuf, ok := os.ReadFile("file.txt")
	//if ok != nil {
	//	panic(ok.Error())
	//}
	//fmt.Println(string(readBuf))

	// 3rd (modern) way of reading file and working with it:
	readBuf, ok := os.Open("file.txt")
	defer readBuf.Close()
	if ok != nil {
		panic(ok.Error())
	}
	textBuf := make([]byte, 100)
	lengthRead, ok := readBuf.Read(textBuf)
	if ok != nil {
		panic(ok.Error())
	}
	fmt.Println("textBuf:", string(textBuf))
	fmt.Println("lengthRead:", lengthRead)

	_, _ = readBuf.Seek(7, 0)
	textBuf = nil
	textBuf = make([]byte, 100)
	//lengthRead, _ = readBuf.Read(textBuf)
	_, _ = readBuf.Seek(0, 0)
	buf, ok := bufio.NewReader(readBuf).Peek(5)
	fmt.Println("textBuf:", string(buf))
}
