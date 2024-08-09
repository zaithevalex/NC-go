package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("time now:", now)
	fmt.Println("year:", now.Year())
	fmt.Println("month:", now.Month())
	fmt.Println("day:", now.Day())
	fmt.Println("hour:", now.Hour())
	fmt.Println("minute:", now.Minute())
	fmt.Println("second:", now.Second())
	fmt.Println("nanosecond:", now.Nanosecond())
	fmt.Println("location:", now.Location())
	fmt.Println("weekday:", now.Weekday())
	fmt.Println()

	birthday := time.Date(2003, 3, 17, 2, 25, 25, 25, time.Local)
	fmt.Println("before case:", now.Before(birthday))
	fmt.Println("after case:", now.After(birthday))
	fmt.Println("equal case:", now.Equal(birthday))
	fmt.Println()

	diffTime := now.Sub(birthday)
	fmt.Println("diff case:", diffTime)
	fmt.Println("diff case hours:", diffTime.Hours())
	fmt.Println("diff case minutes:", diffTime.Minutes())
	fmt.Println("diff case seconds:", diffTime.Seconds())
	fmt.Println("diff case nanoseconds:", diffTime.Nanoseconds())
	fmt.Println()
}
