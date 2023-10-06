package main

import (
	"context"
	"strings"

	pb "github.com/andrew-farries/pgroll-grpc/proto"
	"github.com/xataio/pgroll/pkg/migrations"
	"github.com/xataio/pgroll/pkg/roll"

	"fmt"
)

type PGRollServer struct {
	Roll *roll.Roll
	pb.UnimplementedPGRollServer
}

func (s *PGRollServer) Start(in *pb.StartRequest, stream pb.PGRoll_StartServer) error {
	ctx := context.Background()

	mig, err := migrations.ReadMigration(strings.NewReader(in.GetMigration()))
	if err != nil {
		return err
	}

	cb := func(numRows int64) {
		stream.Send(&pb.StartResponse{Message: fmt.Sprintf("%d rows processed", numRows)})
	}

	if err := s.Roll.Start(ctx, mig, cb); err != nil {
		return err
	}

	stream.Send(&pb.StartResponse{Message: "Done"})

	return nil
}

func (s *PGRollServer) Complete(ctx context.Context, in *pb.CompleteRequest) (*pb.CompleteResponse, error) {
	if err := s.Roll.Complete(ctx); err != nil {
		return nil, err
	}

	return &pb.CompleteResponse{}, nil
}

func (s *PGRollServer) Rollback(ctx context.Context, in *pb.RollbackRequest) (*pb.RollbackResponse, error) {
	if err := s.Roll.Rollback(ctx); err != nil {
		return nil, err
	}

	return &pb.RollbackResponse{}, nil
}
