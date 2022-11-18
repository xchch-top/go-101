package v2

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v2/message"
)

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}
