package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func doGreetManyTimes(c pb.GreeterServiceClient) {
	log.Println("GreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Vova",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
