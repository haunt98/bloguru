package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/haunt98/bloguru/internal/fx/authorfx"
	"go.uber.org/fx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	authorv1 "github.com/haunt98/bloguru/gen/author/v1"
	"google.golang.org/grpc"
)

func main() {
	fx.New(
		authorfx.Module,
		fx.Invoke(initGRPCServer),
	).Run()
}

func initGRPCServer(lc fx.Lifecycle, authorServer authorv1.ServiceServer) error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("failed to listen network: %w", err)
	}

	grpcServer := grpc.NewServer()
	authorv1.RegisterServiceServer(grpcServer, authorServer)
	go func() {
		grpcServer.Serve(lis)
	}()
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	// gRPC-Gateway proxies the requests

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	mux := runtime.NewServeMux()
	if err := authorv1.RegisterServiceHandlerFromEndpoint(context.Background(), mux, ":8080", opts); err != nil {
		return fmt.Errorf("failed to register service handler authorv1: %w", err)
	}

	grpcGatewayServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	go func() {
		grpcGatewayServer.ListenAndServe()
	}()
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return grpcGatewayServer.Shutdown(ctx)
		},
	})

	return nil
}
