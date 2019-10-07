# message.proto
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/hxuanhung/go-grpc-rest-realworld-example \
  --proto_path=third_party \
  --go_out=plugins=grpc:. \
  api/pb/message.proto 

# service.proto
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/hxuanhung/go-grpc-rest-realworld-example \
  --proto_path=third_party \
  --go_out=plugins=grpc:. \
  api/srv/service.proto 

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/hxuanhung/go-grpc-rest-realworld-example \
  --proto_path=third_party \
  --grpc-gateway_out=logtostderr=true:. \
  api/srv/service.proto 

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/hxuanhung/go-grpc-rest-realworld-example \
  --proto_path=third_party \
  --swagger_out=logtostderr=true:api/swagger \
  api/srv/service.proto 