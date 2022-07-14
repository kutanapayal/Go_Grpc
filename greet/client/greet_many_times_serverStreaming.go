package main

import (
	"context"
	pb "go_Grpc/greet/proto"
	"io"
	"log"
)

func DoGreetManyTimes(c pb.GreetServiceClient) {

	log.Printf("DoGreetManyTimes was Invoked.")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Payal",
	})

	if err != nil {
		log.Fatalf("Failed to connect %s\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Printf("End of Stream Results.")
			break
		}
		if err != nil {
			log.Fatalf("Failed to Read %s\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
