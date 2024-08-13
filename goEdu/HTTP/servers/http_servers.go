package main

import (
	"fmt"
	"net/http"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func RESPONSE(w http.ResponseWriter, r *http.Request) {
	_, ok := fmt.Fprintf(w, "RESPONSE")
	ASSERT(ok)
}

func READ_HEADERS(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", RESPONSE)
	http.HandleFunc("/read", READ_HEADERS)
	http.ListenAndServe(":8080", nil)
}
