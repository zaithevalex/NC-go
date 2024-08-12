package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// flag.NewFlagSet(...) registers new command (1st param).
	// flag.NewFlagSet(...).Bool(...) registers new flag for already determined command.
	// For example: newGoCmd1 is new command that was registered by flag.NewFlagSet(...) and
	// command line can detect this command by "cmd1". After this command we can print flags,
	// that determined by flag.newFlagSet(...).Bool(...) and consists from:
	// 1) name of this flag;
	// 2) default meaning of this flag;
	// 3) short description of this flag.
	newGoCmd1 := flag.NewFlagSet("cmd1", flag.ExitOnError)
	newGoCmd1Bool := newGoCmd1.Bool("enable", false, "cmd1 an enable")

	// We can determine another command and flag for it as the same.
	// So, for example, I can determine same command and same flag for this function:
	newGoCmd2 := flag.NewFlagSet("cmd2", flag.ExitOnError)
	newGoCmd2Bool := newGoCmd2.Bool("enable", false, "cmd2 an enable")

	// After it, we can print ONE of the command. After we can print the flag that pinned
	// for this command. Flag has something meaning by default, but it can be changed.
	// WE CAN'T CALL more than 2 new commands in one command line request in one time,
	// because if it happened, last of called command sends to flag.Args().
	switch os.Args[1] {
	case "cmd1":
		newGoCmd1.Parse(os.Args[2:])
		fmt.Println("cmd1 called:")
		fmt.Println("enable:", *newGoCmd1Bool)
		fmt.Println("flag.Args():", newGoCmd1.Args())
	case "cmd2":
		newGoCmd2.Parse(os.Args[2:])
		fmt.Println("cmd2 called:")
		fmt.Println("enable:", *newGoCmd2Bool)
		fmt.Println("flag.Args():", newGoCmd2.Args())
	default:
		fmt.Println("unregistered command")
		os.Exit(1)
	}
}
