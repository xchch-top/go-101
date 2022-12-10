package main

import (
	"context"
	"fmt"
	"gitlab.xchch.top/zhangsai/go-101/training/week17/micro_v4/v1/example/proto/gen"
)

type UserService struct {
	name string
	gen.UnimplementedUserServiceServer
}

func (u *UserService) GetById(ctx context.Context, req *gen.GetByIdReq) (*gen.GetByIdResp, error) {
	fmt.Printf("server %s, get user id: %d \n", u.name, req.Id)
	return &gen.GetByIdResp{
		User: &gen.User{
			Id:     req.Id,
			Status: 123,
		},
	}, nil
}
