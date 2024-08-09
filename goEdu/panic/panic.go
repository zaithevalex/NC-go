package main

import "fmt"

func positiveNumber(number int) error {
	if number < 0 {
		return fmt.Errorf("negative number: %d", number)
	}
	fmt.Println("positive number", number)
	return nil
}

func detectError(ok error) {
	if ok != nil {
		panic(ok)
	}
}

func main() {
	var number1, number2, number3 = 32, -1, 12

	ok := positiveNumber(number1)
	detectError(ok)

	ok = positiveNumber(number2)
	detectError(ok)

	// after error program already won't be finished and number3 won't be print to console:
	ok = positiveNumber(number3)
	detectError(ok)
	panic("END")
}
