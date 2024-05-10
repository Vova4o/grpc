package main

import (
	"context"
	"log"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

func doGreet(c pb.GreeterServiceClient) {
	log.Println("doGreet was called")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Vova",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}

func doSumm(c pb.CalculatorServiceClient) {
	log.Println("doSumm was called")
	res, err := c.Summ(context.Background(), &pb.SummTwoIntsRequest{
		A: 3,
		B: 10,
	})
	if err != nil {
		log.Fatalf("Could not summ: %v\n", err)
	}
	log.Printf("Summ: %d\n", res.Result)
}
