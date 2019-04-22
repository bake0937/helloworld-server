package server

import (
	"context"
	"log"

	pb "proton"
)

type Server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func NewServer() (server *Server) {
	server = &Server{}

	return server
}
