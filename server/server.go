package main

import (
	"log"
	"net"

	"github.com/dojeto/grpc-golang/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Error In Creating Server", err)
	}
	grpcServer := grpc.NewServer()

	s := &Server{}

	proto.RegisterOrderServiceServer(grpcServer, s)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Error In gRPC", err)
	}
}
