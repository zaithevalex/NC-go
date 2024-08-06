package main

import "fmt"

const (
	StateTurnOff = iota
	StateTurnOn
)

var mapState = map[int]string{
	StateTurnOff: "Turn Off",
	StateTurnOn:  "Turn On",
}

func main() {
	state1 := convert(0)
	state2 := convert(1)
	fmt.Println("state1:", state1)
	fmt.Println("state2:", state2)

	stateError := convert(2)
	fmt.Println("stateError:", stateError)
}

func convert(state int) string {
	switch state {
	case StateTurnOff:
		return mapState[StateTurnOff]
	case StateTurnOn:
		return mapState[StateTurnOn]
	default:
		//return ""
		panic("invalid state")
	}
}
