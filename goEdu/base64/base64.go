package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	defaultKey := "default key"
	base64Key := base64.StdEncoding.EncodeToString([]byte(defaultKey))
	fmt.Println("BASE64 KEY:", base64Key)

	base64KeyDecoded, _ := base64.StdEncoding.DecodeString(base64Key)
	fmt.Println("BASE64 KEY ENCODED:", string(base64KeyDecoded))

	urlBase64Key := base64.URLEncoding.EncodeToString([]byte(defaultKey))
	fmt.Println("URL BASE64 KEY:", urlBase64Key)

	urlBase64KeyDecoded, _ := base64.URLEncoding.DecodeString(urlBase64Key)
	fmt.Println("URL BASE64 KEY ENCODED:", string(urlBase64KeyDecoded))
	fmt.Println("BASE64 KEY ENCODED:", base64.StdEncoding.EncodeToString([]byte(urlBase64Key)))

}
