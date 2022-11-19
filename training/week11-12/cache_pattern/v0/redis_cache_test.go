package v0

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/zhangsai/go-101/training/week11-12/cache_pattern/v0/mocks"
	"testing"
	"time"
)

func TestRedisCache_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name       string
		mock       func() redis.Cmdable
		key        string
		val        string
		expiration time.Duration

		wantErr error
	}{
		{
			name: "return OK",
			mock: func() redis.Cmdable {
				res := mocks.NewMockCmdable(ctrl)
				cmd := redis.NewStatusCmd(nil)
				cmd.SetVal("OK")
				res.EXPECT().Set(gomock.Any(), "key1", "value1", time.Minute).
					Return(cmd)
				return res
			},
			key:        "key1",
			val:        "value1",
			expiration: time.Minute,
		},
		{
			name: "timeout",
			mock: func() redis.Cmdable {
				res := mocks.NewMockCmdable(ctrl)
				cmd := redis.NewStatusCmd(nil)
				cmd.SetVal("OK")
				cmd.SetErr(context.DeadlineExceeded)
				res.EXPECT().Set(gomock.Any(), "key1", "value", time.Minute).
					Return(cmd)
				return res
			},
			key:        "key1",
			val:        "value1",
			expiration: time.Minute,
			wantErr:    context.DeadlineExceeded,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmdable := tc.mock()
			client := NewRedisCache(cmdable)
			err := client.Set(context.Background(), tc.key, tc.val, tc.expiration)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
