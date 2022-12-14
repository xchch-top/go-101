package v0

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitClientProxy(t *testing.T) {
	testCases := []struct {
		name string

		service *UserServiceClient
		p       *mockProxy

		wantReq  *Request
		wantResp *GetByIdResp
		wantErr  error
	}{
		{
			name:    "user service",
			service: &UserServiceClient{},
			p: &mockProxy{
				result: []byte(`{"name": "Tom"}`),
			},
			wantReq: &Request{
				ServiceName: "user-service",
				MethodName:  "GetById",
				Data:        []byte(`{"Id":13}`),
			},
			wantResp: &GetByIdResp{
				Name: "Tom",
			},
		},
		{
			name:    "proxy return error",
			service: &UserServiceClient{},
			p: &mockProxy{
				err: errors.New("mock error"),
			},
			wantErr: errors.New("mock error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := InitClientProxy(tc.service, tc.p)
			require.NoError(t, err)

			resp, err := tc.service.GetById(context.Background(), &GetByIdReq{Id: 13})
			if err != nil {
				return
			}
			// 断言p的数据
			assert.Equal(t, tc.wantReq, tc.p.req)
			assert.Equal(t, tc.wantResp, resp)
		})
	}
}

type mockProxy struct {
	req    *Request
	result []byte
	err    error
}

func (m *mockProxy) Invoke(ctx context.Context, req *Request) (*Response, error) {
	m.req = req
	return &Response{
		Data: m.result,
	}, m.err
}

type UserServiceClient struct {
	GetById func(ctx context.Context, req *GetByIdReq) (*GetByIdResp, error)
}

func (u *UserServiceClient) Name() string {
	return "user-service"
}

type GetByIdReq struct {
	Id int64
}

type GetByIdResp struct {
	Name string `json:"name"`
}
