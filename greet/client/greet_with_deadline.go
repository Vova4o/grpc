package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreeterServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Vova",
	}

	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Timeout was exceeded")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Printf("A non gRPC error: %v\n", err)
		}
		return
	}
	log.Printf("GreetWithDeadLine: %s\n", res.Result)
}
