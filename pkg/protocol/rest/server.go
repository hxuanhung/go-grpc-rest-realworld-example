package rest

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/api/v1"
	"github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	middleware "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/protocol/rest/middleware"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, httpPort, grpcPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterAcreProtectionHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}
	srv := &http.Server{
		Addr: ":" + httpPort, // ":" is required
		// add handler with middleware
		Handler: middleware.AddRequestID(
			middleware.AddLogger(logger.Log, mux)),
	}

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
