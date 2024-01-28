package main

import (
	"log"

	"github.com/ffajarpratama/grpc-sm/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:5052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect %v\n", err)
	}

	log.Printf("dialed to %s\n", addr)

	defer conn.Close()

	c := proto.NewCalculatorServiceClient(conn)
	// DoSum(c)
	DoPrimes(c)
}
