package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)

	var sum int64 = 0
	for index := range len(nums) {
		sum += int64(nums[index])
	}
	fmt.Println("sum:", sum)

	// for-each cycle style
	for num := range nums {
		fmt.Println("num:", num)
	}

	// intermediate for-each cycle with index of element
	for i, num := range nums {
		fmt.Printf("i: %d | num: %d\n", i, num)
	}

	m := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("m:", m)

	for k, v := range m {
		fmt.Printf("%d: %s\n", k, v)
		//fmt.Println(k, v)
	}

	// print keys. If only alone variable in a cycle, then will print keys from map
	for k := range m {
		fmt.Printf("key: %d\n", k)
	}

	// for + string:
	for index, unicode := range "Hello, world!" {
		fmt.Println(index, unicode)
	}
	fmt.Println()

	// also u can do next
	s := "Hello, world!"
	for index, unicode := range s {
		fmt.Println(index, unicode)
	}
}
