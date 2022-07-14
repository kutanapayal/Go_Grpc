package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"io"
	"log"
)

func Primedivions(c pb.PrimeServiceClient) {
	log.Printf("Client function invoked.")
	stream, err := c.Primedivions(context.Background(), &pb.PrimeRequest{
		Num: 120,
	})

	if err != nil {
		log.Fatalf("Failed to Connect : %s\n", err)
	}
	for {
		res, msg := stream.Recv()

		if msg == io.EOF {
			log.Printf("Thanks For Using Our Service.")
			break
		}
		if msg != nil {
			log.Fatalf("Sorry, Failed To Read Try Again Later...%s\n", msg)
		}
		log.Printf("%v", res.Result)
	}

}
