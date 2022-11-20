//go:build e2e

package v4

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v4/example/rpc/proto/gen"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v4/example/rpc/server"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v4/serialize/proto"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(":8081")
	require.NoError(t, err)

	us := &server.UserServiceProto{}
	err = c.InitService(us)
	c.serializer = proto.Serializer{}
	require.NoError(t, err)

	resp, err := us.GetById(context.Background(), &gen.GetByIdReq{Id: 100})
	require.NoError(t, err)
	log.Println(resp)
}
