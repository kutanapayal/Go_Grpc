package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"io"
	"log"
	"time"
)

func findMax(c pb.MaxServiceClient) {
	log.Printf("Client Invoked..")

	reqs := []*pb.MaxRequest{
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 6},
		{Num: 2},
		{Num: 20},
	}

	stream, err := c.FindMax(context.Background())

	if err != nil {
		log.Fatalf("Failed to connect ...%s\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("%v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				break
			}
			if err != nil {
				log.Fatalf("Failed to Read ...%s\n", err)
			}
			log.Printf("Max Out Of this :%v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
