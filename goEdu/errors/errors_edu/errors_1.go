package main

import "fmt"

type ArgError struct {
	number  int
	message string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.number, e.message)
}

func divide(nominator int, denominator int) error {
	fmt.Printf("nominator: %d, denominator: %d\n", nominator, denominator)
	if denominator == 0 {
		return &ArgError{1, "divide by zero"}
	}
	return nil
}

func main() {
	var nominator = 1
	//var denominator = 0
	//err := divide(nominator, denominator)
	//fmt.Println(err)

	err1 := ArgError{
		number:  nominator,
		message: "divide by zero",
	}

	fmt.Println(err1)  // <=> err1, because doesn't exist (e ArgError) Error() string {...}
	fmt.Println(&err1) // <=> (*(&err1)).Error()
	var err2 = &err1
	fmt.Println(&err2)

}
