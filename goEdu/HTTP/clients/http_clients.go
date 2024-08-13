package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	response, ok := http.Get("https://google.com")
	ASSERT(ok)
	defer response.Body.Close()
	fmt.Println("response status:", response.Status)
	//fmt.Println("response headers:", response.Header)
	//fmt.Println("response body:", response.Body)

	scanner := bufio.NewScanner(response.Body)
	for index := 0; scanner.Scan(); index++ {
		fmt.Println(scanner.Text())
	}
	ASSERT(scanner.Err())
}
