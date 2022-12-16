package ratelimit

import "context"

type Limiter interface {
	Acquire(ctx context.Context, req any) (any, error)
	Release(resp any, err error)
}

type Guardian interface {
	Allow(ctx context.Context, req any) (func(), error)
	OnRejection(ctx context.Context, req any) (any, error)
}
