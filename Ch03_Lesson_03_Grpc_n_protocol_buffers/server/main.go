package main

import (
	"context"
	pb "github.com/dNaszta/ms_go/Ch03_Lesson_03_Grpc_n_protocol_buffers/fibonacci"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
	num = 4967295
)

type server struct {
	cache [num]uint64
	pb.UnimplementedFibonacciServer
}

func (s *server) mustEmbedUnimplementedFibonacciServer() {}

func (s *server) Calculate(ctx context.Context, in *pb.FibonacciRequest) (*pb.FibonacciReply, error) {
	timeStart := time.Now()
	result := s.CalculateFibonacci(in.Number)

	return &pb.FibonacciReply{
		Result:         result,
		ProcessingTime: uint64(time.Since(timeStart)),
	}, nil
}

func (s *server) CalculateFibonacci(num uint32) uint64 {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	res := s.cache[num]
	if res != 0 {
		return res
	}

	res = s.CalculateFibonacci(num-2) + s.CalculateFibonacci(num-1)
	s.cache[num] = res
	return res
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFibonacciServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}