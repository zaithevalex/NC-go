package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MatchString("^[A-Za-z0-9]*$", "general"))

	// save pattern for using after:
	reg1, _ := regexp.Compile("^[A-Za-z]*$") // ^ - start of string; $ - end of string;
	reg2, _ := regexp.Compile("[A-Za-z]*")
	reg3, _ := regexp.Compile("[A-Za-z]+")

	fmt.Println("reg.MatchString:", reg1.MatchString("aAa!"))
	fmt.Println("reg.FindString:", reg2.FindString("aAa! AKDnkFnf asfasf asdas"))
	fmt.Println("reg.FindStringIndex:", reg3.FindStringIndex("!!!! AKDnkFnf asfasf asdas"))
	fmt.Println("reg.FindAllString:", reg3.FindAllString("!!!! AKDnkFnf asfasf asdas \t", -1))

	// same, that is regexp.Compile, but it doesn't return 2nd arg (error)
	reg4 := regexp.MustCompile("[A-Za-z]+")
	fmt.Println("reg4:", reg4)

	// change meaning with determined regexp function:
	fmt.Println("reg.ReplaceAllString:", reg4.ReplaceAllString("!!!! AKDnkFnf asfasf asdas \t", "1_2"))
}
