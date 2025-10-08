package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"products/src/pb/products"
	"products/src/repository"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
	repo *repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.repo.Create(*product)
	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	productList, err := s.repo.FindAll()
	return &productList, err
}

func main() {
	fmt.Println("Starting grpc server")
	srv := server{}

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	products.RegisterProductServiceServer(s, &srv)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
