package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("--- listBlog was invoked ---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs RPC: %v", err)
	}
	for {
		blog, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		log.Println(blog)
	}
}
