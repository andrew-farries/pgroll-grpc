package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/andrew-farries/pgroll-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port        = 50051
	postgresURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

func main() {
	ctx := context.Background()

	roll, err := NewRoll(ctx)
	if err != nil {
		log.Fatalf("failed to start pgroll: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPGRollServer(grpcServer, &PGRollServer{Roll: roll})

	fmt.Printf("Server listening on port %d...", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
