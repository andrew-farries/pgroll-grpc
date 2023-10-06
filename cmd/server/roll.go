package main

import (
	"context"

	"github.com/xataio/pgroll/pkg/roll"
	"github.com/xataio/pgroll/pkg/state"
)

func NewRoll(ctx context.Context) (*roll.Roll, error) {
	const lockTimeoutMs = 500

	state, err := state.New(ctx, postgresURL, "pgroll")
	if err != nil {
		return nil, err
	}

	roll, err := roll.New(ctx, postgresURL, "public", lockTimeoutMs, state)
	if err != nil {
		return nil, err
	}

	if err := roll.Init(ctx); err != nil {
		return nil, err
	}

	return roll, nil
}
