package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/110y/gae-go-grpc/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	grpcPort = 5000
	httpPort = 8080
)

type server struct {
}

func (s *server) Echo(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	return &pb.Message{
		Message: fmt.Sprintf("HELLO: %s", req.Message),
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		fmt.Printf("failed to listen on port: %d", grpcPort)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterApiServer(s, &server{})

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		s.Serve(lis)
	}()

	go func() {
		runHTTPServer(ctx)
	}()

	select {
	case sig := <-sigChan:
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			s.GracefulStop()
		}
	}
}

func runHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterApiHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcPort), opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}
