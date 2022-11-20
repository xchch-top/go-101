package main

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v4/v0/example/rpc/proto/gen"
)

// UserServiceProto 用来测试 protobuf 协议
type UserServiceProto struct {
}

func (u *UserServiceProto) GetById(ctx context.Context, req *gen.GetByIdReq) (*gen.GetByIdResp, error) {
	return &gen.GetByIdResp{
		User: &gen.User{
			Id: 123,
		},
	}, nil
}

func (u *UserServiceProto) Name() string {
	return "user-service-proto"
}
