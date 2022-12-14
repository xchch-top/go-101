package ratelimit

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Limiter interface {
	Acquire(ctx context.Context, req any) (any, error)
	Release(resp any, err error)
}

type Guardian interface {
	Allow(ctx context.Context, req any) (func(), error)
	OnRejection(ctx context.Context, req any) (any, error)
}

type rejectStrategy func(ctx context.Context, info *grpc.UnaryServerInfo,
	req any, handler grpc.UnaryHandler) (any, error)

var defaultRejectStrategy = func(ctx context.Context, info *grpc.UnaryServerInfo,
	req any, handler grpc.UnaryHandler) (any, error) {
	return nil, status.Errorf(codes.ResourceExhausted, "触发限流 %s", info.FullMethod)
}
