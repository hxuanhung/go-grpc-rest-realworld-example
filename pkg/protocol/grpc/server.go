package grpcmiddleware

import (
	"context"
	"fmt"
	"net"

	srv "github.com/hxuanhung/go-grpc-rest-realworld-example/api/srv"
	"github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/logger"
	middleware "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/protocol/grpc/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, server srv.GreeterServer, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		logger.Log.Fatal("failed to listen", zap.String("reason", err.Error()))
	}
	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = testdata.Path("server1.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = testdata.Path("server1.key")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		logger.Log.Fatal("Failed to generate credentials", zap.String("reason", err.Error()))
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	grpcServer := grpc.NewServer(opts...)
	srv.RegisterGreeterServer(grpcServer, server)

	logger.Log.Info("starting gRPC server...")
	grpcServer.Serve(lis)
}
