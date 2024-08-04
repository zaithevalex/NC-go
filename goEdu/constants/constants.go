package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func main() {
	fmt.Println(s, reflect.TypeOf(s))

	const number1 int64 = 100
	fmt.Println(number1, reflect.TypeOf(number1))

	const number2 int32 = -256
	fmt.Println(math.Abs(float64(number2)), reflect.TypeOf(number2))

	var number3 = -int16(number2)
	fmt.Println(number3, reflect.TypeOf(number3))
	fmt.Println("sin of number3: ", math.Sin(float64(number3)))
}
