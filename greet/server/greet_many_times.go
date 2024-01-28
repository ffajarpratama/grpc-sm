package main

import (
	"fmt"

	pb "github.com/ffajarpratama/grpc-sm/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("in: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
