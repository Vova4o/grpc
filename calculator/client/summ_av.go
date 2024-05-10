package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
)

func doSendManyNUmbers(c pb.PrimeNumberServiceClient) {
	log.Println("doSendManyNUmbers was invoked")

	numbers := []*pb.GetPrimeNumbers{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.AverageOfSumm(context.Background())
	if err != nil {
		log.Fatalf("Error while calling AverageOfSumm: %v\n", err)
	}

	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		stream.Send(number)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from AverageOfSumm: %v\n", err)
	}
	log.Printf("AverageOfSumm: %v\n", res.Result)
}
