package main

import (
	desc "ClientServerProject/protobuf/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func ASSERT(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	connection, err := grpc.Dial(fmt.Sprintf("localhost:%d", 8080), grpc.WithTransportCredentials(insecure.NewCredentials()))
	ASSERT(err)
	defer connection.Close()
	client := desc.NewDescriptionServiceClient(connection)
	//_ = desc.NewDescriptionServiceClient(connection)
	call(client)
}

func call(client desc.DescriptionServiceClient) {
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.GetDescription(context, nil)
	ASSERT(err)
	log.Printf("%s\n", reply.Description)
}
