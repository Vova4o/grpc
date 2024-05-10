package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func doLongGreet(c pb.GreeterServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Vova",
		},
		{
			FirstName: "Vova2",
		},
		{
			FirstName: "Vova3",
		},
		{
			FirstName: "Vova4",
		},
		{
			FirstName: "Vova5",
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}
	log.Printf("LongGreet: %s\n", res.Result)
}
