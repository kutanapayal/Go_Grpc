package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Dosqrt(c pb.SqrtServiceClient, n int32) {
	log.Println("Sqrt client Invoked")

	res, err := c.FindSqrt(context.Background(), &pb.SqrtRequest{Num: n})

	if err != nil {

		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error Message :: %s\n", e.Message())
			log.Printf("Error Code :: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("Number is Invalide.")
			}
		} else {
			log.Fatalf("Non Grpc Error... %s\n", err)
		}
	}

	log.Printf("Sqrt of N:%v is:%v", n, res.Result)
}
