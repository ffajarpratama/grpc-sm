package main

import (
	"fmt"

	"github.com/ffajarpratama/grpc-sm/calculator/proto"
)

func (s *Server) Primes(in *proto.PrimeRequest, stream proto.CalculatorService_PrimesServer) error {
	fmt.Printf("in: %v\n", in)

	var k int32 = 2
	for in.Number > 1 {
		if in.Number%k == 0 {
			stream.Send(&proto.PrimeResponse{Result: k})
			in.Number /= k
		} else {
			k++
		}
	}

	return nil
}
