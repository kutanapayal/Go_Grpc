package main

import (
	"context"
	"fmt"
	pb "go_Grpc/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) CreatBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {

	log.Print("Create Blog was Invoked.")
	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	fmt.Println("Date: ", data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v\n", err),
		)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	fmt.Println("ID: ", id)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot Convert ID",
		)
	}

	return &pb.BlogId{
		Id: id.Hex(),
	}, nil

}
