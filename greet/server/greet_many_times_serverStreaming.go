package main

import (
	"fmt"
	pb "go_Grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("GreetManyTimes function was involved with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s : %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil

}
