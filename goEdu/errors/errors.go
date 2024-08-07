package main

import (
	"fmt"
)

type ArgError struct {
	number  int
	message string
}

// var NegativeNumbersError = errors.New("can't work with negative numbers.")
var NegativeNumbersError = fmt.Errorf("can't work with negative numbers.")

func square(height, width int) (int, error) {
	if height < 0 || width < 0 {
		return -1, NegativeNumbersError
	}
	return height * width, nil
}

func checkNumber(number int) error {
	if number < 0 {
		return NegativeNumbersError
	}
	return nil
}

func (e ArgError) Error() string {
	e.number = 1000
	return fmt.Sprintf("%d - %s", e.number, e.message)
}

func checkNumber2(number int) (int, error) {
	if number < 0 {
		return -1, &ArgError{number, "can't work with negative numbers."}
	}
	return number, nil
}

func ERROR_INFO(argError ArgError) {
	fmt.Printf("error number: %d | error message: %s\n", argError.number, argError.message)
}

func main() {
	//fmt.Println(square(10, 10))
	//fmt.Println(square(-1, 1))
	//fmt.Println(checkNumber(1))
	//fmt.Println(checkNumber(-1))
	//
	//_, err1 := checkNumber2(-1)
	//_, err2 := checkNumber2(-2)
	//
	//fmt.Println("########################")
	//fmt.Println("err1:", err1)
	//fmt.Println("err2:", err2)
	//fmt.Println(errors.Is(err1, err2))
	//fmt.Println(errors.As(err1, &err2))

	//_, err := checkNumber2(-1)
	//var argError *ArgError
	//fmt.Println("error:", reflect.TypeOf(err))
	//fmt.Println("error x:", reflect.TypeOf(&argError))
	//fmt.Println(errors.As(err, &argError))
	//err = StrErr{}
	//var res *StrErr = &StrErr{}
	//ok := errors.As(err, res)
	//fmt.Println(ok)

	//fmt.Println("err1.(*ArgError):", err1.(*ArgError))
	//fmt.Println("err1:", err1)

	err := ArgError{number: 40, message: "error"}
	fmt.Println("#1 err:", err.number, err)
	_ = err.Error()
	fmt.Println("#2 err:", err.number, err)
}

//type StrErr []string
//
//func (StrErr) Error() string {
//	return ""
//}
