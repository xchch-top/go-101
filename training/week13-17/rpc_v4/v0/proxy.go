package v0

import (
	"context"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v4/v0/message"
)

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}
