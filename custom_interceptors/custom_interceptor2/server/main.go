package main

import (
	desc "EDU/protobuf"
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type DescriptionServer struct {
	desc.DescriptionServiceServer
}

var runner chan bool

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println(err.Error())
	}

	//grpcServer := grpc.NewServer()
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor))
	grpcServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	desc.RegisterDescriptionServiceServer(grpcServer, &DescriptionServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err.Error())
	}

	<-runner
}

func (descServer *DescriptionServer) GetDescription(ctx context.Context, request *desc.DescriptionRequest) (*desc.DescriptionReply, error) {
	log.Println("###LOG MESSAGE:", request.GetId())
	return &desc.DescriptionReply{
		Description: &desc.Description{
			Id:      request.GetId(),
			Content: fmt.Sprintf("message #%d", request.GetId()),
		},
	}, nil
}

func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//log.Println("UnaryServerInterceptor PRE:", info.FullMethod)
	//m, err := handler(ctx, req)
	//if err != nil {
	//	log.Println(err.Error())
	//	return nil, err
	//}

	//log.Println("UnaryServerInterceptor POST:", info.FullMethod)
	//return m, nil

	log.Println("#####INTECEPTOR !!!")
	return otelgrpc.UnaryServerInterceptor()(ctx, req, info, handler)
}
