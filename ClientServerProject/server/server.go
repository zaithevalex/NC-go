package main

import (
	desc "ClientServerProject/protobuf/protobuf"
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type DescriptionItem struct {
	id      int64  `bson:"id"`
	content string `bson:"content"`
}

type DescriptionServer struct {
	desc.DescriptionServiceServer
}

func ASSERT(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	ASSERT(err)
	//var opts []grpc.ServerOption
	//grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	desc.RegisterDescriptionServiceServer(grpcServer, &DescriptionServer{})
	err = grpcServer.Serve(listener)
	ASSERT(err)
}

func (descriptionServer *DescriptionServer) GetDescription(context context.Context, request *desc.DescriptionRequest) (*desc.DescriptionReply, error) {
	return &desc.DescriptionReply{
		Description: &desc.Description{
			Id:      1,
			Content: "1st message",
		},
	}, nil
}
