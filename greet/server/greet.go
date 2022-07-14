package main

import (
	"context"
	pb "go_Grpc/greet/proto"
	"log"
)

func (s *Server) Greet(c context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet Function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
