.PHONY: build run

build:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pgroll.proto

run: build
	go run ./cmd/server
