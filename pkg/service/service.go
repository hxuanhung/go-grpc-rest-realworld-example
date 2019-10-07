package service

import (
	"context"

	"github.com/hxuanhung/go-grpc-rest-realworld-example/api/pb"
)

// Server struct
type Server struct{}

// HealthCheck implements helloworld.ServiceServer
func (s *Server) HealthCheck(ctx context.Context, in *pb.GetHealthCheckRequest) (*pb.GetHealthCheckResponse, error) {
	// log.Printf("Received: %v", in.GetName())
	return &pb.GetHealthCheckResponse{Status: "Good!"}, nil
}
