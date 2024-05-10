package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func (s *Server) LongGreet(sream pb.GreeterService_LongGreetServer) error {
	log.Println("LongGreet was invoked")

	res := ""
	for {
		req, err := sream.Recv()
		if err == io.EOF {
			return sream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
