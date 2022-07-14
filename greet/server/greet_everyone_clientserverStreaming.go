package main

import (
	pb "go_Grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Server Was Invoked.")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error While Reading Client...")
		}
		res := "Greeting to " + req.FirstName
		stream.Send(&pb.GreetResponse{Result: res})
	}
}
