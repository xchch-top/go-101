//go:build e2e

package v3

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(":8081")
	require.NoError(t, err)

	us := &UserServiceClient{}
	err = c.InitService(us)
	require.NoError(t, err)

	resp, err := us.GetById(context.Background(), &GetByIdReq{Id: 100})
	require.NoError(t, err)
	log.Println(resp)
}
