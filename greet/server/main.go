package main

import (
	"log"
	"net"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreeterServiceServer
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	opts := []grpc.ServerOption{}
	tls := true // change that to false if you don't want TLS

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to load certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pb.RegisterGreeterServiceServer(s, &Server{})
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
