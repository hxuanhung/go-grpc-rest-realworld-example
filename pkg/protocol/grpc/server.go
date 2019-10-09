package grpcmiddleware

import (
	"context"
	"fmt"
	"net"

	v1 "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/api/v1"
	"github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/logger"
	middleware "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/protocol/grpc/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish service
func RunServer(ctx context.Context, server v1.AcreProtectionServer, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		logger.Log.Fatal("failed to listen", zap.String("reason", err.Error()))
	}
	var opts []grpc.ServerOption

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	grpcServer := grpc.NewServer(opts...)
	v1.RegisterAcreProtectionServer(grpcServer, server)

	logger.Log.Info("starting gRPC server...")
	grpcServer.Serve(lis)
}
