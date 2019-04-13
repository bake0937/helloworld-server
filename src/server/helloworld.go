package server

import (
	"context"
	"log"

	pb "proton"
	//‎⁨Macintosh HD⁩ ▸ ⁨ユーザ⁩ ▸ ⁨komoriryuuta⁩ ▸ ⁨project⁩ ▸ ⁨src⁩ ▸ ⁨github.com⁩ ▸ ⁨ryutakomori⁩ ▸ ⁨protoc-gen-docker⁩ ▸ ⁨src⁩ ▸ ⁨go⁩
	//pb "../../protoc-gen-docker⁩/src⁩/go⁩"
	//pb "github.com/ryutakomori/protoc-gen-docker⁩/src/go"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
