package main

import (
	"context"
	"log"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- updateBlog was invoked ---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Vova",
		Title:    "My first blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was updated: %s\n", id)
}
