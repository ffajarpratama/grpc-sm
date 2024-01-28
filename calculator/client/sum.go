package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ffajarpratama/grpc-sm/calculator/proto"
)

func DoSum(c proto.CalculatorServiceClient) {
	log.Println("do sum called")

	res, err := c.Sum(context.Background(), &proto.SumRequest{FirstNum: 3, SecondNum: 10})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %+v\n", res)
}
