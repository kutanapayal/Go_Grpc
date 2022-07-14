package main

import (
	"context"
	pb "go_Grpc/blog/proto"
	"log"
)

func CreatBlog(c pb.BlogServiceClient) string {
	log.Printf("Create Blog CLient Was Invoked.")

	blog := &pb.Blog{
		AuthorId: "Mox",
		Title:    "My First Go Blog",
		Content:  "Go is the most dengerious language",
	}

	res, err := c.CreatBlog(context.Background(), blog)

	if err != nil {
		log.Printf("Unexpected Error %s\n", err)
	}

	log.Printf("Blog Created With ID : %v\n", res)

	return res.Id
}
