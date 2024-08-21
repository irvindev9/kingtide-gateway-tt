package main

import (
	"log"

	"context"

	"google.golang.org/grpc"

	"client-api/pkg/api"
)

// func createProduct(product *api.Products) string {

// 	var conn *grpc.ClientConn
// 	conn, err := grpc.NewClient(":5001", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %s", err)
// 	}
// 	defer conn.Close()

// 	c := api.NewProductsServiceClient(conn)

// 	response, err := c.CreateProducts(context.Background(), &api.Products{
// 		Name:        product.Name,
// 		Price:       product.Price,
// 		Quantity:    product.Quantity,
// 		Description: product.Description,
// 	})
// 	if err != nil {
// 		log.Fatalf("Error when calling CreateProducts: %s", err)
// 	}
// 	log.Printf("Response from server: %s", response.String())

// 	return response.String()

// }

func getProducts() string {

	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewProductsServiceClient(conn)

	response, err := c.CreateProducts(context.Background(), &api.Products{})
	if err != nil {
		log.Fatalf("Error when calling CreateProducts: %s", err)
	}
	log.Printf("Response from server: %s", response.String())

	return response.String()

}
