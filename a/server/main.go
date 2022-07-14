package main

import (
	pb "go_Grpc/a/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type server struct {
	pb.AServiceServer
	pb.PrimeServiceServer
	pb.AvgServiceServer
	pb.MaxServiceServer
	pb.SqrtServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", addr)
	}
	s := grpc.NewServer()
	pb.RegisterAServiceServer(s, &server{})
	pb.RegisterPrimeServiceServer(s, &server{})
	pb.RegisterAvgServiceServer(s, &server{})
	pb.RegisterMaxServiceServer(s, &server{})
	pb.RegisterSqrtServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to server : ", err)
	}
}
