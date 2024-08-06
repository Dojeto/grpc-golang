package main

import (
	"context"
	"fmt"

	"github.com/dojeto/grpc-golang/proto"
)

type Server struct {
	proto.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, order *proto.OrderRequest) (respose *proto.OrderResponse, err error) {
	fmt.Println("Order Requested From Client :", order.UserID, "For Order : ", order.ProductID)
	// Perform Operation On Database or whatever you want
	return &proto.OrderResponse{
		OrderId: 1,
		Status:  "You Order is Done",
	}, nil
}
