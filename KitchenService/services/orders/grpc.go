package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	// Create a tcp connection
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen :%s", err.Error())
	}

	// and send it to grpcSever.Serve()

	grpcSever := grpc.NewServer()
	// Register our grpc services
	log.Println("starting grpc server on : ", s.addr)
	return grpcSever.Serve(lis)
}
