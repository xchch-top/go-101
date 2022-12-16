package v2

import (
	"context"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v2/message"
)

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}
