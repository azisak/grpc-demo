package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"example.com/demo-grpc/server"
	"example.com/demo-grpc/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8082", "gRPC server endpoint")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := ":8082"
	httpAddr := ":8081"
	lisGrpc, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// GRPC Setup
	s := grpc.NewServer()
	user_service.RegisterUserServiceServer(s, server.NewController())

	// Http Setup
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = user_service.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Serving GRPC
	log.Info("Serving gRPC on https://", grpcAddr)
	go func() {
		log.Fatal(s.Serve(lisGrpc))
	}()

	// Serving Rest
	log.Info("Serving Rest on https://", httpAddr)
	http.ListenAndServe(httpAddr, mux)
}
