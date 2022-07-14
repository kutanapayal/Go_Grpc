package main

import (
	"context"
	pb "go_Grpc/blog/proto"
	"log"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("DeleteBlog was Invoked.")

	req := &pb.BlogId{Id: id}
	res, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Printf("Failed to Delete: %v", err)
	}

	log.Printf("Blog has been Deleted. %v\n", res)
}
