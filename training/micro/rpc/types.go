package rpc

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/micro/rpc/message"
)

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}

type Service interface {
	ServiceName() string
}
