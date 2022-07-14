package main

import (
	pb "go_Grpc/blog/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addrs string = "localhost:50051"

func main() {

	tls := true // changed when it's needed
	opts := []grpc.DialOption{}

	if tls {
		cert := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(cert, "") // "" to not ovverride the address

		if err != nil {
			log.Fatalf("Failed Load creadentials %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addrs, opts...)

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Connect()
	//...
	c := pb.NewBlogServiceClient(conn)

	id := CreatBlog(c)

	DeleteBlog(c, id)

	// ReadBlog(c, id)

	// UpdateBLog(c, id)

	// ReadBlog(c, id)

	// ListBlog(c)
}
