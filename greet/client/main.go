package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/vova4o/grpc-go-study/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("failed to load CA certificates: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	}

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterServiceClient(conn)
	// b := pb.NewCalculatorServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)

	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 1*time.Second)

	// doSumm(b)
}
