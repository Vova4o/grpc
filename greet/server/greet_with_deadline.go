package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadLine(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadLine function was invoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
