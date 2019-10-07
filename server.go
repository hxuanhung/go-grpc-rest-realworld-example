package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"

	"github.com/hxuanhung/go-grpc-rest-realworld-example/api/pb"
	"github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/logger"
	grpc "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/protocol/grpc"
	rest "github.com/hxuanhung/go-grpc-rest-realworld-example/pkg/protocol/rest"
)

var (
	tls           = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile      = flag.String("cert_file", "", "The TLS cert file")
	keyFile       = flag.String("key_file", "", "The TLS key file")
	jsonDBFile    = flag.String("json_db_file", "", "A json file containing a list of features")
	port          = flag.String("port", "localhost:10000", "The gRPC server port")
	httpPort      = flag.String("httpPort", "8081", "The rest server port")
	logLevel      = flag.Int("logLevel", -1, "Global log level")
	logTimeFormat = flag.String("logTimeFormat", "2006-01-02T15:04:05.999999999Z07:00", "Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) HealthCheck(ctx context.Context, in *pb.GetHealthCheckRequest) (*pb.GetHealthCheckResponse, error) {
	// log.Printf("Received: %v", in.GetName())
	return &pb.GetHealthCheckResponse{Status: "Good!"}, nil
}

func main() {
	ctx := context.Background()
	// initialize logger
	if err := logger.Init(*logLevel, *logTimeFormat); err != nil {
		fmt.Errorf("failed to initialize logger: %v", err)
		return
	}

	flag.Parse()
	defer glog.Flush()

	// Goroutines: A goroutine is a lightweight thread managed by the Go runtime.
	go func() {
		grpc.RunServer(ctx, &server{}, *port)
	}()

	if err := rest.RunServer(ctx, *httpPort, *port); err != nil {
		glog.Fatal(err)
	}
}
