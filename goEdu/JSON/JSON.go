package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// !!! IMPORTANT: naming must be with upper case !!!
type Human struct {
	Name     string
	Age      uint8
	Features []string
}

func main() {
	aloneBoolValueJson, _ := json.Marshal(true)
	// 1st parameter in Println after string is byte slice of content JSON struct:
	fmt.Println("aloneBoolValueJson:", aloneBoolValueJson, string(aloneBoolValueJson))
	// if we use only alone variable, then structure of json.Marshal will analog

	slice := []int{1, 2, 3, 4, 5}
	sliceJson, _ := json.Marshal(slice)
	fmt.Println("sliceJson:", string(sliceJson))

	Map := map[string]int{"first": 1, "second": 2}
	mapJson, _ := json.Marshal(Map)
	fmt.Println("mapJson:", string(mapJson))

	human1 := &Human{
		Name:     "Aleksandr",
		Age:      21,
		Features: []string{"GO", "JAVA"},
	}

	human2 := Human{
		"Dima",
		31,
		[]string{"GO", "JAVA"},
	}

	humanJson1, _ := json.Marshal(human1)
	fmt.Println("humanJson1:", string(humanJson1))

	humanJson2, _ := json.Marshal(human2)
	stringHuman2JSON := string(humanJson2)
	fmt.Println("humanJson2:", stringHuman2JSON)

	// decoding already existence JSON-string:
	//var humanJson3 map[string]interface{}
	var humanJson3 map[string]any
	if err := json.Unmarshal([]byte(stringHuman2JSON), &humanJson3); err != nil {
		panic(err.Error())
	}
	fmt.Println()
	fmt.Println("humanJson3:", humanJson3)
	fmt.Println()
	humanJson3Name := humanJson3["Name"]
	humanJson3Age := humanJson3["Age"]
	humanJson3Features := humanJson3["Features"]
	fmt.Printf("humanJson3Name:\ntype: %v\ncontent: %v\n\n", reflect.TypeOf(humanJson3Name), humanJson3Name)
	fmt.Printf("humanJson3Age:\ntype: %v\ncontent: %v\n\n", reflect.TypeOf(humanJson3Age), humanJson3Age)
	fmt.Printf("humanJson3Features:\ntype: %v\ncontent: %v\n\n", reflect.TypeOf(humanJson3Features), humanJson3Features)

	// translation to console, but allows to translate in a file or HTTP-response, for example:
	fmt.Println("TRANSLATION TO CONSOLE:")
	encoding := json.NewEncoder(os.Stdout)
	if ok := encoding.Encode(humanJson3); ok != nil {
		panic(ok.Error())
	}
}
