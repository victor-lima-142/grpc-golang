package main

import (
	"context"
	"fmt"
	"log"
	"products/src/pb/products"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	defer conn.Close()

	client := products.NewProductServiceClient(conn)
	productsBeforeInsert, err := client.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatal("failed to find all products:", err)
	}
	fmt.Println("Products before insert:", productsBeforeInsert)

	client.Create(context.Background(), &products.Product{
		Name:        "Notebook",
		Description: "Notebook of banana brand",
		Price:       300.0,
		Quantity:    10,
	})
	client.Create(context.Background(), &products.Product{
		Name:        "Phone",
		Description: "Phone of banana brand",
		Price:       100.0,
		Quantity:    10,
	})
	client.FindAll(context.Background(), &products.Product{})

	productsAfterInsert, err := client.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatal("failed to find all products:", err)
	}
	fmt.Println("Products before insert:", productsAfterInsert)
}
