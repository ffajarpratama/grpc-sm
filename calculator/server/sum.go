package main

import (
	"context"
	"fmt"

	"github.com/ffajarpratama/grpc-sm/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	fmt.Printf("in: %v\n", in)

	res := &proto.SumResponse{
		Result: in.FirstNum + in.SecondNum,
	}

	return res, nil
}
