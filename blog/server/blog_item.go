package main

import (
	pb "go_Grpc/blog/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"id, omitempty" json:"id,omitempty"`
	AuthorID string             `bson:"author_id" json:"author_id"`
	Title    string             `bson:"title" json:"title"`
	Content  string             `bson:"content" json:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		// Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
