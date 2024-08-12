package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf := scanner.Text()
		fmt.Println("scanned now:", buf)
	}
}
