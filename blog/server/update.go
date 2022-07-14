package main

import (
	"context"
	pb "go_Grpc/blog/proto"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	log.Printf("UpdateBLog Was Invoked.")
	id, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	set := bson.M{"$set": data}
	filter := bson.M{"_id": id}

	res, err := collection.UpdateOne(ctx, filter, set)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Coud not Update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find blog with the ID provided",
		)
	}

	return &emptypb.Empty{}, nil
}
