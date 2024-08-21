package main

import (
	"log"

	"context"

	"google.golang.org/grpc"

	"client-api/pkg/api"
)

func createOrder(order *api.Orders) string {

	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewOrdersServiceClient(conn)

	response, err := c.CreateOrders(context.Background(), &api.Orders{
		ClientName: order.ClientName,
		SaleDate:   order.SaleDate,
		OrderSales: order.OrderSales,
	})
	if err != nil {
		log.Fatalf("Error when calling CreateOrders: %s", err)
	}
	log.Printf("Response from server: %s", response.ResponseMessage)

	return response.ResponseMessage

}
