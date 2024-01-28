package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ffajarpratama/grpc-sm/calculator/proto"
)

func DoPrimes(c proto.CalculatorServiceClient) {
	log.Println("do primes called")

	req := &proto.PrimeRequest{
		Number: 1080,
	}

	stream, err := c.Primes(context.Background(), req)
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
