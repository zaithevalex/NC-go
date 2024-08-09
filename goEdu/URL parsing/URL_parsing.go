package main

import (
	"fmt"
	"net"
	"net/url"
)

func MapInfo(Map url.Values) {
	for key, value := range Map {
		fmt.Printf("key: %s | value: %s\n", key, value[0])
	}
}

func main() {
	link := "https://user:pass@host.com:5432/path?k=v&l=u#f"
	URL, ok := url.Parse(link)
	if ok != nil {
		panic(ok)
	}

	fmt.Println("scheme:", URL.Scheme)
	fmt.Println("user info:", url.User)
	fmt.Println("username:", URL.User.Username())

	password, _ := URL.User.Password()
	fmt.Println("password:", password)
	fmt.Println("host:", URL.Host)

	host, port, _ := net.SplitHostPort(URL.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)

	fmt.Println("path:", URL.Path)
	fmt.Println("fragment:", URL.Fragment)
	fmt.Println("RawQuery:", URL.RawQuery)

	Map, _ := url.ParseQuery(URL.RawQuery)
	MapInfo(Map)
}