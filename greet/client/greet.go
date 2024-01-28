package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ffajarpratama/grpc-sm/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet called")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "kimi"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %+v\n", res)
}
