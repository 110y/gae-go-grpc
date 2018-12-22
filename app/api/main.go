package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/110y/gae-go-grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Echo(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	return &pb.Message{
		Message: fmt.Sprintf("HELLO: %s", req.Message),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5000))
	if err != nil {
		// TODO:
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterApiServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		// TODO:
		panic(err)
	}
}
