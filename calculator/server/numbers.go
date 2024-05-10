package main

import (
	"log"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
)

func (s *Server) Numbers(in *pb.GetPrimeNumbers, stream pb.PrimeNumberService_NumbersServer) error {
	log.Printf("Numbers function was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)


	for number > 1 {
		if number % divisor == 0 {
			res := &pb.GetPrimeNumbersResponse{
				Result: divisor,
			}
			if err := stream.Send(res); err != nil {
				log.Fatalf("Error while sending stream: %v\n", err)
			}
			number = number / divisor
		} else {
			divisor++
		}
	}

	return nil
}
