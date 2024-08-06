package main

import (
	"context"
	"log"

	pb "github.com/dojeto/grpc-golang/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewOrderServiceClient(conn)
	resp, _ := client.CreateOrder(context.Background(), &pb.OrderRequest{
		ProductID: 32,
		UserID:    2,
	})

	log.Println(resp)
}
