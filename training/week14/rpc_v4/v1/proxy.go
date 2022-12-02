package v1

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v4/v1/message"
)

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}
