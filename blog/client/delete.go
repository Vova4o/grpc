package main

import (
	"context"
	"log"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- deleteBlog was invoked ---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was deleted: %s\n", id)
}
