DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# gRPC and gRPC-gateway
echo "Generate gRPC stub in Go and generate reverse-proxy..."
protoc -I/usr/local/include -I. \
  -I $GOPATH/src \
  -I $DIR/api/proto/v1 \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:pkg/api/v1 \
  --grpc-gateway_out=logtostderr=true:pkg/api/v1 \
  protection-service.proto 
echo "done!"

# swagger
# Note that the proto_path must be an exact prefix of the .proto file names
echo "Generate Swagger files..."
protoc -I/usr/local/include -I. \
  -I $GOPATH/src \
  -I $DIR/api/proto/v1 \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:api/swagger/v1 \
  protection-service.proto 
echo "done!"
