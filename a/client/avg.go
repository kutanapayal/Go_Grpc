package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"
	"time"
)

func DoAvg(c pb.AvgServiceClient) {
	reqs := []*pb.AvgRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}

	stream, err := c.FindAvg(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect %s\n", err)
	}

	for i, req := range reqs {
		log.Printf("Number%v : %v", i+1, req.Num)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Failed to Read.")
	}

	log.Printf("Avg : %v", res.Avg)
}
