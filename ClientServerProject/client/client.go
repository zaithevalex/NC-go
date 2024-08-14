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
	//call(client, nil)

	for index := 0; index < 10; index++ {
		var id = desc.DescriptionRequest{
			Id: int64(index),
		}
		call(client, &id)
		time.Sleep(2 * time.Second)
	}
}

func call(client desc.DescriptionServiceClient, in *desc.DescriptionRequest) {
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.GetDescription(context, in)
	ASSERT(err)
	log.Printf("%s\n", reply.Description)
}
