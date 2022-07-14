package main

import (
	"context"
	pb "go_Grpc/blog/proto"
	"log"
)

func ReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was Invoked.")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Failed to Read: %v", err)
	}

	log.Printf("Blog Was Read :%v\n", res)

	return res
}
