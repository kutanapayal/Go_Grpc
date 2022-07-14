package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"
	"math"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Findabs(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Server Invoked...%v", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(
				codes.Canceled, "Client Cancel the request",
			)
		}
		time.Sleep(1 * time.Second)

	}

	return &pb.SqrtResponse{
		Result: float32(math.Abs(float64(in.Num))),
	}, nil
}
