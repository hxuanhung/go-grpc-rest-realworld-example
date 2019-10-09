.PHONY: build
build:
	./protoc-gen.sh

.PHONY: dev
dev:
	realize start

.PHONY: start-server
start-server:
	go run ./cmd/server/server.go