package api

import (
	"log"

	"context"
)

type Server struct {
}

func (s *Server) CreateOrders(ctx context.Context, in *OrdersResponse) (*OrdersResponse, error) {
	log.Printf("Receive message body from client: %s", in.ResponseMessage)
	return &OrdersResponse{ResponseMessage: in.ResponseMessage}, nil
}

func (s *Server) GetProducts(ctx context.Context, in *ProductsResponse) (*ProductsResponse, error) {
	log.Printf("Receive message body from client: %s", in.String())
	return &ProductsResponse{}, nil
}
