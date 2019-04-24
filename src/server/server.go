package server

import (
	"context"
	"log"
	"model"
	pb "proton"
)

type Server struct {
	model *model.Model
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func NewServer(m *model.Model) (server *Server) {
	server = &Server{
		model: m,
	}

	return server
}
