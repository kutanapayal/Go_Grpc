package main

import (
	pb "go_Grpc/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", addr)
	}

	log.Printf("Listening on %s\n", addr)

	tls := true
	opts := []grpc.ServerOption{}

	if tls {
		cert := "ssl/server.crt"
		keyfile := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(cert, keyfile)
		if err != nil {
			log.Fatalf("Failed Loading Certificates %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Server: %s\n", err)
	}

}
