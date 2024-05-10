package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
)

func doPrimeNumber(c pb.PrimeNumberServiceClient) {
	log.Println("doPrimeNumber was invoked")

	req := &pb.GetPrimeNumbers{
		Number: 120,
	}

	stream, err := c.Numbers(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Numbers: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("GetPrimeNumbers: %d\n", res.Result)
	}
}
