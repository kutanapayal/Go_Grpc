package main

import (
	pb "go_Grpc/a/proto"
	"io"
	"log"
)

func (s *server) FindMax(stream pb.MaxService_FindMaxServer) error {
	log.Print("Server was invoked ...")
	var max int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to Read...%s\n", err)
		}
		if max < req.Num {
			max = req.Num
			stream.Send(&pb.MaxResponse{Result: max})
		}
	}
}
