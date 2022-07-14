package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"
)

func (s *server) A(ctx context.Context, in *pb.ARequest) (*pb.AResponse, error) {
	log.Printf("Server Invoked...%v", in)
	return &pb.AResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
