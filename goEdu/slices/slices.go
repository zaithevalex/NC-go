package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	// nil <=> null pointer or zero-length
	fmt.Println("slice:", s == nil, len(s) == 0)

	// initializing already existed slice by 3 elements
	s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Printf("#1 length: %d\ncapacity: %d\n", len(s), cap(s))

	// if it was array, we couldn't add adding elem. Slices decide this problem:
	s = append(s, "d")
	fmt.Printf("#2 length: %d\ncapacity: %d\n", len(s), cap(s))
	s = append(s, "e", "f")
	fmt.Printf("#3 length: %d\ncapacity: %d\n", len(s), cap(s))
	// if capacity = length, slice will increase x2 (current capacity = previous capacity * 2)
	s = append(s, "g")
	fmt.Printf("#4 length: %d\ncapacity: %d\n", len(s), cap(s))

	var c = make([]string, 10)
	copy(c, s)
	fmt.Println(c)
	fmt.Printf("length: %d\ncapacity: %d\n", len(c), cap(c))

	// 5th element doesn't copy !
	d := s[2:5]
	fmt.Printf("length: %d\ncapacity: %d\n", len(d), cap(d))

	// also we can copy before 5th element, not including 5th element
	d = s[:5]
	fmt.Println(d)
	fmt.Printf("length: %d\ncapacity: %d\n", len(d), cap(d))

	d = s[2:]
	fmt.Println(d)
	fmt.Printf("length: %d\ncapacity: %d\n", len(d), cap(d))

	var e = d
	fmt.Printf("e length: %d\ncapacity: %d\n", len(d), cap(d))
	fmt.Printf("e length: %d\ncapacity: %d\n", len(e), cap(e))

	t1 := []string{"a", "b", "c", "d"}
	t2 := []string{"a", "b", "c"}

	// slices are equals if each of element of 1st slice = each of element 2nd slice
	fmt.Println("Equal(t1, t2):", slices.Equal(t1, t2))
	t1 = t2
	// now t1 contains elements from t2 slice (copy)
	fmt.Println("Equal(t1, t2):", slices.Equal(t1, t2))
	fmt.Printf("links: %p %p\n", &t1, &t2)

	t3 := []string{"a", "b", "c", "d"}
	t4 := []string{"a", "b", "c", "d"}
	fmt.Printf("links: %p %p\n", &t3, &t4)
	fmt.Println(slices.Equal(t3, t4))
	t3 = t4
	fmt.Printf("links: %p %p\n", &t3, &t4)
	fmt.Println(slices.Equal(t3, t4))

	//double slices:
	t5 := make([][]int, 3)
	fmt.Println(t5)
	for i := 0; i < len(t5); i++ {
		fmt.Printf("t5[%d] len: %d; cap: %d\n", i, len(t5[i]), cap(t5[i]))
		for j := 0; j < i+1; j++ {
			t5[i] = append(t5[i], j)
		}
		fmt.Printf("t5[%d] len: %d; cap: %d\n", i, len(t5[i]), cap(t5[i]))
	}
	fmt.Println(t5, "\n\n\n")

	t6 := make([][]int, 3)
	fmt.Println(t6)
	for i := range len(t6) {
		// here already I must to allocate memory for solo-dimensional slice
		t6[i] = make([]int, i+1)
		fmt.Printf("t6[%d] len: %d; cap: %d\n", i, len(t6[i]), cap(t6[i]))
		for j := range i + 1 {
			t6[i][j] = j
		}
		fmt.Printf("t6[%d] len: %d; cap: %d\n", i, len(t6[i]), cap(t6[i]))
	}
	fmt.Println(t6)
}
