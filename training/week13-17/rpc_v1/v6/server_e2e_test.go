//go:build e2e

package v6

import (
	"context"
	"testing"
)

func TestServer_Start(t *testing.T) {
	s := NewServer()
	_ = s.Register(&UserServiceServer{})
	s.Start(":8081")
}

type UserServiceServer struct {
}

func (u *UserServiceServer) Name() string {
	return "user-service"
}

func (u *UserServiceServer) GetById(ctx context.Context, req *GetByIdReq) (*GetByIdResp, error) {
	return &GetByIdResp{
		Name: "Tom",
	}, nil
}
