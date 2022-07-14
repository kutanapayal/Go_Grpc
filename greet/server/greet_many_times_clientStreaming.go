package main

import (
	"fmt"
	pb "go_Grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreeting(stream pb.GreetService_LongGreetingServer) error {
	log.Println("Server was Invoked.")

	res := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("Failed to Read.")
		}

		res += fmt.Sprintf("Greeting %s\n", req.FirstName)

	}

}
