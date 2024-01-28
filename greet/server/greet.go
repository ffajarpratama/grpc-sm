package main

import (
	"context"
	"fmt"

	pb "github.com/ffajarpratama/grpc-sm/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("in: %v\n", in)

	res := &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}

	return res, nil
}
