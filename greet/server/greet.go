package main

import (
	"context"
	"log"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) Summ(ctx context.Context, in *pb.SummTwoIntsRequest) (*pb.SummTwoIntsResponse, error) {
	log.Printf("Summ function was invoked with %v\n", in)
	return &pb.SummTwoIntsResponse{
		Result: in.A + in.B,
	}, nil
}