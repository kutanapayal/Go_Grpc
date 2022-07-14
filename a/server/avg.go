package main

import (
	pb "go_Grpc/a/proto"
	"io"
	"log"
)

func (s *server) FindAvg(stream pb.AvgService_FindAvgServer) error {
	log.Println("Server Invoked..")
	avg := 0.0
	count, sum := 0.0, 0.0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			avg = sum / count
			return stream.SendAndClose(&pb.AvgResponse{Avg: float32(avg)})
		}
		if err != nil {
			log.Fatal("Error While REading.")
		}

		sum += float64(req.Num)
		count++

	}
}
