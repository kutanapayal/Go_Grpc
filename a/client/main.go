package main

import (
	pb "go_Grpc/a/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addrs string = "0.0.0.0:50052"

func main() {

	conn, err := grpc.Dial(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect :%s", err)
	}

	defer conn.Connect()

	// c := pb.NewAServiceClient(conn)
	// DoSum(c)
	// s := pb.NewPrimeServiceClient(conn)
	// Primedivions(s)
	// c := pb.NewAvgServiceClient(conn)
	// DoAvg(c)
	// c := pb.NewMaxServiceClient(conn)
	// findMax(c)
	c := pb.NewSqrtServiceClient(conn)
	// Dosqrt(c, -6)
	findabs(c, 1*time.Second)
}
