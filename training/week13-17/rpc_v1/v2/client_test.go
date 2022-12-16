package v2

import (
	"context"
	"fmt"
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
			p := &mockProxy{
				result: []byte(`{"name": "Tom"}`),
			}
			err := InitClientProxy(tc.service, p)
			assert.Equal(t, tc.wantErr, err)

			resp, err := tc.service.GetById(context.Background(), &GetByIdReq{Id: 13})
			require.Nil(t, err)
			// 断言p的数据
			assert.Equal(t, &Request{
				ServiceName: "user-service",
				MethodName:  "GetById",
				Arg:         &GetByIdReq{Id: 13},
			}, p.req)
			assert.Equal(t, &GetByIdResp{
				Name: "Tom",
			}, resp)
			fmt.Println(resp)
		})
	}
}

type mockProxy struct {
	req    *Request
	result []byte
}

func (m *mockProxy) Invoke(ctx context.Context, req *Request) (*Response, error) {
	m.req = req
	return &Response{
		Data: m.result,
	}, nil
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
	Name string `json:"name"`
}
