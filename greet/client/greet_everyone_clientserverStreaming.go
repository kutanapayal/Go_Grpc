package main

import (
	"context"
	pb "go_Grpc/greet/proto"
	"io"
	"log"
	"time"
)

func GreetEveryone(c pb.GreetServiceClient) {
	log.Printf("Client Invoked...")

	reqs := []*pb.GreetRequest{
		{FirstName: "payal"},
		{FirstName: "Mok"},
		{FirstName: "Lovve"},
	}

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error Whilw connecting ...%s\n", err)
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending :", req.FirstName)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error Whilw Reading ...%s\n", err)
			}
			log.Printf("Greeting :%s\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc

}
