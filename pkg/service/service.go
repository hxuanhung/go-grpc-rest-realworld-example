package service

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/hxuanhung/go-grpc-rest-realworld-example/api/pb"
	"github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/logger"
	"go.uber.org/zap"
)

// Server struct
type Server struct {
	savedUsers []*pb.User // read-only after initialized
}

// HealthCheck implements srv.AcreServer
func (s *Server) HealthCheck(ctx context.Context, in *pb.GetHealthCheckRequest) (*pb.GetHealthCheckResponse, error) {
	// log.Printf("Received: %v", in.GetName())
	return &pb.GetHealthCheckResponse{Status: "Good!"}, nil
}

// UserLogin implements srv.AcreServer
func (s *Server) UserLogin(ctx context.Context, input *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	// log.Printf("Received: %v", in.GetName())
	for _, user := range s.savedUsers {
		if user.Email == input.Email {
			return &pb.UserLoginResponse{Token: "Generate a token here!"}, nil
		}
	}
	// No feature was found, return an unnamed feature
	// return &pb.Feature{Location: point}, nil
	return &pb.UserLoginResponse{Token: "NOT GOOD!"}, nil
}

// loadUsers loads users from a JSON file.
func (s *Server) loadUsers(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			logger.Log.Fatal("Failed to load default features", zap.String("reason", err.Error()))
		}
	} else {
		data = exampleData
	}
	if err := json.Unmarshal(data, &s.savedUsers); err != nil {
		logger.Log.Fatal("Failed to load default features: %v", zap.String("reason", err.Error()))
	}
}

// exampleData is a copy of testdata/user_db.json. It's to avoid
// specifying file path with `go run`.
var exampleData = []byte(`[{
	"email": "hung",
    "passwordHash": "123"
}]`)
