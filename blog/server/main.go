package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	collection *mongo.Collection
	addr       string = "0.0.0.0:50051"
)

type Server struct {
	pb.BlogServiceServer
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// outdated code
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	// if err != nil {
	// 	log.Fatalf("failed to create client: %v", err)
	// }
	// err = client.Connect(context.Background())
	// if err != nil {
	// 	log.Fatalf("failed to connect: %v", err)
	// }
	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
