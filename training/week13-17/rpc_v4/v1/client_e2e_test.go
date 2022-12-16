//go:build e2e

package v1

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(":8081")
	require.NoError(t, err)

	us := &UserServiceClient{}
	err = c.InitService(us)
	require.NoError(t, err)

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*3)
	defer cancelFunc()
	resp, err := us.GetById(ctx, &GetByIdReq{Id: 100})
	require.NoError(t, err)
	log.Println(resp)
}
