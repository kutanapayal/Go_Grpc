package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) FindSqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Server Invoked...%v", in)

	if in.Num < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			"Invalide Argument...",
		)
	}

	return &pb.SqrtResponse{
		Result: float32(math.Sqrt(float64(in.Num))),
	}, nil
}
