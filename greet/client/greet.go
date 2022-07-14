package main

import (
	"context"
	pb "go_Grpc/greet/proto"
	"log"
)

func DoGreeet(c pb.GreetServiceClient) {
	log.Println("DoGreet Invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Payal",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting : %s\n", res.Result)
}
