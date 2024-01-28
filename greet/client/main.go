package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ffajarpratama/grpc-sm/greet/proto"
)

var addr string = "0.0.0.0:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect %v\n", err)
	}

	log.Printf("dialed to %s\n", addr)

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	doGreetManyTimes(c)
}
