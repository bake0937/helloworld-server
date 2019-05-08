package server

import (
	"context"
	"log"
	pb "proton"
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
