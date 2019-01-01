package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/datastore"
	pb "github.com/110y/gae-go-grpc/app/api/proto"
	"github.com/110y/gae-go-grpc/internal/lib/appengine/grpc/interceptor"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

const (
	grpcPort = 5000
	httpPort = 8080
)

var client *datastore.Client

type user struct {
	ID   string `datastore:"-"`
	Name string
}

func Execute() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := loadEnvironmentVariables(); err != nil {
		return err
	}

	err := initializeDatastoreClient(ctx, env.GcpProjectID)
	if err != nil {
		return fmt.Errorf("failed to create a datastore client: %+v", err)
	}

	if env.EnableStackdriverTrace {
		flush, err := setupStackdriverTrace(ctx, env.GcpProjectID)
		if err != nil {
			return fmt.Errorf("failed to set up stackdriver trace: %+v", err)
		}
		defer flush()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen on port: %d", grpcPort)
	}

	opts := []grpc.ServerOption{
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryServerInterceptors()...)),
	}

	s := grpc.NewServer(opts...)
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

	return nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Printf("exit: %+v", err)
		os.Exit(1)
	}
}

func runHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			EmitDefaults: true,
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", grpcPort), opts...)
	if err != nil {
		return err
	}

	apiClient := pb.NewApiClient(conn)
	pb.RegisterApiHandlerClient(ctx, mux, apiClient)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}

func initializeDatastoreClient(ctx context.Context, project string) error {
	c, err := datastore.NewClient(ctx, project)
	if err != nil {
		return err
	}

	client = c

	return nil
}

func unaryServerInterceptors() []grpc.UnaryServerInterceptor {
	var interceptors []grpc.UnaryServerInterceptor

	if env.AppEngineNamespace != "" {
		interceptors = append(
			interceptors,
			interceptor.AppEngineNamespacingUnaryServerInterceptor(env.AppEngineNamespace),
		)
	}

	return interceptors
}
