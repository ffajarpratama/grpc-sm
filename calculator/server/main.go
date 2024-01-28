package main

import (
	"log"
	"net"

	"github.com/ffajarpratama/grpc-sm/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5052"

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%v\n", err)
	}
}
