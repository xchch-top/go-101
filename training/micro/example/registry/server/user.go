package main

import (
	"context"
	"fmt"
	"gitlab.xchch.top/golang-group/go-101/training/micro/example/proto/gen"
)

type UserService struct {
	gen.UnimplementedUserServiceServer
}

func (u *UserService) GetById(ctx context.Context, req *gen.GetByIdReq) (*gen.GetByIdResp, error) {
	fmt.Printf("user id: %d", req.Id)
	return &gen.GetByIdResp{
		User: &gen.User{
			Id:     req.Id,
			Status: 123,
		},
	}, nil
}
