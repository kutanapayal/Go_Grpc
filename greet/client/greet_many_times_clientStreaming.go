package main

import (
	"context"
	pb "go_Grpc/greet/proto"
	"log"
	"time"
)

func LongGreeting(c pb.GreetServiceClient) {
	log.Println("Client Streaming Invoked.")

	reqs := []*pb.GreetRequest{
		{FirstName: "payal"},
		{FirstName: "Mok"},
		{FirstName: "Love"},
	}

	stream, err := c.LongGreeting(context.Background())

	if err != nil {
		log.Printf("Failed to connect with server.")
	}

	for _, req := range reqs {
		log.Printf("sending req :%s\n", req)
		stream.Send(req)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Failed to Read response from server.%s\n", err)
	}

	log.Printf("%s\n", res.Result)

}
