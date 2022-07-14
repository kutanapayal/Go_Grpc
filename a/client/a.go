package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"
)

func DoSum(c pb.AServiceClient) {
	log.Println("Sum Invoked")

	res, err := c.A(context.Background(), &pb.ARequest{
		Num1: 10,
		Num2: 3,
	})

	if err != nil {
		log.Fatalf("Failed to call DoSum")
	}

	log.Printf("Sum :%v", res.Result)
}
