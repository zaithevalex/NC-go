package main

import (
	desc "EDU/protobuf"
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	//connection, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	}
	connection, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Println(err.Error())
	}
	defer connection.Close()

	client := desc.NewDescriptionServiceClient(connection)
	id := desc.DescriptionRequest{
		Id: int64(12),
	}
	call(client, &id)
	time.Sleep(2 * time.Second)
}

func call(client desc.DescriptionServiceClient, in *desc.DescriptionRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.GetDescription(ctx, in)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(reply.Description)
}
