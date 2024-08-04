package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Print(i, "\t")
		i++
	}

	fmt.Println()
	var index = int64(1)
	for index <= 10 {
		fmt.Print(index, "\t")
		index++
	}

	fmt.Println()
	for index := range 3 { // here's index is local value and it's not 'index' that was initialized earlier;
		fmt.Print(index, "\t")
	}

	fmt.Println()
	fmt.Println(index)

	//infinity cycle and BREAK:
	for {
		fmt.Print("1", "\t")
		break
	}

	// CONTINUE
	fmt.Println()
	j := int64(1)
	for {
		j++
		if j == 4 {
			continue
		}
		if j == 11 {
			break
		}

		fmt.Print(j-1, "\t")
	}
}
