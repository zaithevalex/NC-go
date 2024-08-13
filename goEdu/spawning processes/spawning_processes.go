package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	dateCmd := exec.Command("date")
	fmt.Println("dateCmd:", dateCmd)
	dateOut, err := dateCmd.Output()
	ASSERT(err)
	fmt.Println("> date")
	fmt.Println("dateOut:", string(dateOut))

	//grepCmd := exec.Command("grep", "hello")
	////fmt.Println("grepCmd:", grepCmd)
	//grepIn, _ := grepCmd.StdinPipe()
	//grepOut, _ := grepCmd.StdoutPipe()
	//////fmt.Println("grepIn:", grepIn)
	//////fmt.Println("grepOut:", grepOut)
	//_ = grepCmd.Start()
	//_, _ = grepIn.Write([]byte("message"))
	//_ = grepIn.Close()
	//grepBytes, _ := ioutil.ReadAll(grepOut)
	//_ = grepCmd.Wait()
	//fmt.Println("> grep hello")
	//fmt.Println("grepBytes:", string(grepBytes))

	grepCmd := exec.Command("grep", "message")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	_ = grepCmd.Start()
	_, _ = grepIn.Write([]byte("message"))
	_ = grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	_ = grepCmd.Wait()
	fmt.Println("> grep message")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	ASSERT(err)
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
