package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ffajarpratama/grpc-sm/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do greet many times called")

	req := &pb.GreetRequest{
		FirstName: "kimi",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error-reading-stream: %v\n", err)
		}

		fmt.Printf("msg: %v\n", msg)
	}
}
