package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	authorv1 "github.com/haunt98/bloguru/gen/author/v1"
	"github.com/haunt98/bloguru/internal/author"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	authorServer := author.NewServer()
	grpcServer := grpc.NewServer()
	authorv1.RegisterServiceServer(grpcServer, authorServer)
	go func() {
		grpcServer.Serve(lis)
	}()

	// gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()
	if err := authorv1.RegisterServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
