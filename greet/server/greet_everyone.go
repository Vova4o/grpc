package main

import (
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreeterService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		res := "Hello " + req.FirstName + "! "
		sendErr := stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v\n", sendErr)
		}
	}
}
