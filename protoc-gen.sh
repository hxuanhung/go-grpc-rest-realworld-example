# message.proto
protoc -I/usr/local/include -I. \
  -I $GOPATH/src \
  -I third_party \
  --go_out=plugins=grpc:. \
  api/pb/*.proto 

# gRPC and gRPC-gateway
protoc -I/usr/local/include -I. \
  -I $GOPATH/src \
  -I third_party \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. \
  api/srv/*.proto 

# swagger
# Note that the proto_path must be an exact prefix of the .proto file names
protoc -I/usr/local/include -I. \
  -I $GOPATH/src \
  -I $GOPATH/src/github.com/hxuanhung/go-grpc-rest-realworld-example/api/srv \
  -I third_party \
  --proto_path=third_party \
  --swagger_out=logtostderr=true:api/swagger \
  service.proto 