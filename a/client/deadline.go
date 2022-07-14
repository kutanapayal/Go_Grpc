package main

import (
	"context"
	pb "go_Grpc/a/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func findabs(c pb.SqrtServiceClient, timeout time.Duration) {
	log.Println("client was invoked.")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.Findabs(ctx, &pb.SqrtRequest{Num: -10})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {

			if e.Code() == codes.DeadlineExceeded {
				log.Println("DEadline WAs EXceeded")
				return
			} else {
				log.Fatalf("Unexpected Grpc Error %s", err)
			}

		} else {
			log.Fatalf("Non Grpc error %s", err)
		}
	}

	log.Printf("Abs of Num is: %v", res.Result)

}
