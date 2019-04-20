package server

import (
	"context"
	"fmt"
	"log"

	pb "proton"
)

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("aaaaaaaaaa")
	log.Printf("Received: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
