package main

import (
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
)

func (s *Server) AverageOfSumm(stream pb.PrimeNumberService_AverageOfSummServer) error {
	log.Println("AverageOfSumm function was invoked")

	sum := 0
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GetAverageSummNumbersResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		sum += int(req.Number)
		count++
	}
}
