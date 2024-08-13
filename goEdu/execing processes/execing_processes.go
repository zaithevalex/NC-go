package main

import (
	"os"
	"os/exec"
	"syscall"
)

func ASSERT(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	binary, lookErr := exec.LookPath("ls")
	ASSERT(lookErr)
	//args := []string{"ls"}
	args := []string{"lsls"}
	environment := os.Environ()
	execErr := syscall.Exec(binary, args, environment)
	ASSERT(execErr)
}
