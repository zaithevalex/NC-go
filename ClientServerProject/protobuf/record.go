package main

import (
	desc "ClientServerProject/protobuf/protobuf"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	user1 := desc.User{
		Login:    "zaithevalex@gmail.com",
		Password: "password",
		Age:      21,
	}

	// proto serialization:
	dataPb, _ := proto.Marshal(&user1)
	dataJSON, _ := json.Marshal(&user1)
	fmt.Println("Protobuf length:", len(dataPb))
	fmt.Println("JSON length:", len(dataJSON))

	// proto deserialization:
	userDecode := &desc.User{}
	_ = proto.Unmarshal(dataPb, userDecode)
	fmt.Println("User decode:", userDecode)
}
