package main

import (
	pb "go_Grpc/a/proto"
	"log"
)

func (s *server) Primedivions(in *pb.PrimeRequest, stream pb.PrimeService_PrimedivionsServer) error {

	log.Printf("Server was invoked with %v\n", in)

	var k uint32 = 2
	var n uint32 = in.Num
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n = n / k
		} else {
			k = k + 1
		}

	}
	return nil
}
