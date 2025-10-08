package main

import (
	"calc/src/pb/calc"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error on new client. ", err)
	}
	defer conn.Close()

	client := calc.NewCalcServiceClient(conn)
	stream, err := client.Calc(context.Background())
	if err != nil {
		log.Fatalln("error on get channel to stream. ", err)
	}

	numbersToSend := []int32{10, 2, 5, 3}
	for _, v := range numbersToSend {
		if err = stream.Send(&calc.Input{
			Value: v,
		}); err != nil {
			log.Fatalln("error on send to stream. ", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error on receive from stream. ", err)
	}

	log.Println(response)
}
