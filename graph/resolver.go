package graph

import (
	"context"
	"log/slog"

	"github.com/uptrace/bun"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ctx    *context.Context
	logger *slog.Logger
	db     *bun.DB
}

func NewResolver(ctx *context.Context, logger *slog.Logger, db *bun.DB) *Resolver {
	return &Resolver{
		ctx:    ctx,
		logger: logger,
		db:     db,
	}
}
