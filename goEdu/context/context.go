package main

import (
	"fmt"
	"net/http"
	"time"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HELLO(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: Hello handler started")
	defer fmt.Println("server: Hello handler ended")

	for {
		select {
		case <-time.After(2 * time.Second):
			_, _ = fmt.Fprintf(w, "hello\n")
		case <-ctx.Done():
			err := ctx.Err()
			//fmt.Println("TRIGGER!!!")
			fmt.Println("server:", err.Error())
			interalError := http.StatusInternalServerError
			http.Error(w, err.Error(), interalError)
			//default:
			//	time.Sleep(200 * time.Millisecond)
			//	_, _ = fmt.Fprintf(w, "default")
		}
	}
}

func main() {
	http.HandleFunc("/hello", HELLO)
	http.ListenAndServe(":8080", nil)
}
