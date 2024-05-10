package main

import (
	"context"
	"log"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--- readBlog was invoked ---")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was read: %s\n", res)

	return res
}
