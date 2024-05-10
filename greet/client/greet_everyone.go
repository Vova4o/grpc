package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func doGreetEveryone(c pb.GreeterServiceClient) {
	log.Println("doGreetEveryone was invoked")

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

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()
	

	<-waitc
}
