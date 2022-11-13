package v1

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitClientProxy(t *testing.T) {
	testCases := []struct {
		name    string
		service *UserService
		wantErr error
	}{
		{
			name:    "user service",
			service: &UserService{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := InitClientProxy(tc.service)
			assert.Equal(t, tc.wantErr, err)
			// 缺乏校验手段
			_, err = tc.service.GetById(context.Background(), &GetByIdReq{Id: 13})
			require.NotNil(t, err)
			str := err.Error()
			req := Request{Arg: &GetByIdReq{}}
			err = json.Unmarshal([]byte(str), &req)
			require.Nil(t, err)
			assert.Equal(t, Request{
				ServiceName: "user-service",
				MethodName:  "GetById",
				Arg: &GetByIdReq{
					Id: 13,
				},
			}, req)
		})
	}
}

type UserService struct {
	GetById func(ctx context.Context, req *GetByIdReq) (*GetByIdResp, error)
}

func (u *UserService) Name() string {
	return "user-service"
}

type GetByIdReq struct {
	Id int64
}

type GetByIdResp struct {
}
