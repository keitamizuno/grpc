package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	result := in.FirstNum + in.SecondNum

	return &pb.SumResponse{
		Result: result,
	}, nil
}
