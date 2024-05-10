package main

import (
	"log"
	"net"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PrimeNumberServiceServer
	pb.MaxNumberServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()

	pb.RegisterPrimeNumberServiceServer(s, &Server{})
	pb.RegisterMaxNumberServiceServer(s, &Server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
