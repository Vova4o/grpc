package main

import (
	"context"
	"log"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--- createBlog was invoked ---")

	blog := &pb.Blog{
		AuthorId: "Vova",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
