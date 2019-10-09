package service

import (
	"context"
	"encoding/json"
	"io/ioutil"

	v1 "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/api/v1"
	"github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/logger"
	"go.uber.org/zap"
)

// Server struct
type Server struct {
	savedUsers []*v1.UserLoginRequest_User // read-only after initialized
}

// HealthCheck implements srv.AcreServer
func (s *Server) HealthCheck(ctx context.Context, in *v1.GetHealthCheckRequest) (*v1.GetHealthCheckResponse, error) {
	// log.Printf("Received: %v", in.GetName())
	return &v1.GetHealthCheckResponse{Status: "Good!"}, nil
}

// UserLogin implements srv.AcreServer
func (s *Server) UserLogin(ctx context.Context, req *v1.UserLoginRequest) (*v1.UserLoginResponse, error) {
	logger.Log.Debug("Received:", zap.Any("input", req))
	for _, user := range s.savedUsers {
		if user.Email == req.User.Email && user.Password == req.User.Password {
			return &v1.UserLoginResponse{Token: "Generate a token here!", Email: req.User.Email}, nil
		}
	}
	// No feature was found, return an unnamed feature
	// return &v1.Feature{Location: point}, nil
	return &v1.UserLoginResponse{Token: "NOT GOOD!"}, nil
}

// LoadUsers loads users from a JSON file.
func (s *Server) LoadUsers(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			logger.Log.Fatal("Failed to load default features", zap.String("reason", err.Error()))
		}
	} else {
		logger.Log.Fatal("Failed to load default users: DB file is empty")
	}
	logger.Log.Debug("DB", zap.ByteString("users", data))
	if err := json.Unmarshal(data, &s.savedUsers); err != nil {
		logger.Log.Fatal("Failed to load default users: %v", zap.String("reason", err.Error()))
	}
}
