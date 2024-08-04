package main

import (
	"fmt"
	"maps"
)

func main() {
	var m = make(map[string]int)
	fmt.Println(m)
	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3
	fmt.Println(m)

	var meaning1 = m["k1"]
	var meaning2 = m["k2"]
	fmt.Println("meaning1:", meaning1)
	fmt.Println("meaning2:", meaning2)
	fmt.Println("map length:", len(m))

	delete(m, "k2")
	fmt.Println("m:", m)
	fmt.Println("map length:", len(m))

	clear(m)
	fmt.Println("m:", m)
	fmt.Println("map length:", len(m))
	fmt.Println("not exists:", m["k2"])

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	m["k2"] = 2
	_, meaning3 := m["k3"]
	fmt.Println("meaning3:", meaning3)

	var n = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("n:", n)

	// compares each elem of each map
	var n1 = map[string]int{"a": 1, "b": 2, "c": 3}
	var n2 = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println("n == n1:", maps.Equal(n, n1))
	fmt.Println("n == n2:", maps.Equal(n, n2))
}
