package main

import (
	"context"
	pb "go_Grpc/blog/proto"
	"log"
)

func UpdateBLog(c pb.BlogServiceClient, id string) {

	log.Println("UpdateBlog WAs Invoked.")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Payal",
		Title:    "A new Python BLog",
		Content:  "Python is really robust....",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error While Updating: %v\n", err)
	}

	log.Println("Blog Was Updated.")

}
