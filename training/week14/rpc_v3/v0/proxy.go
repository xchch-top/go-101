package v0

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v3/v0/message"
)

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}
