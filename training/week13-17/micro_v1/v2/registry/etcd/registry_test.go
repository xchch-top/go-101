package etcd

import (
	"github.com/stretchr/testify/assert"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
)

func TestRegistry_Subscribe(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func() clientV3.Watcher
		wantErr error
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &Registry{
				client: &clientV3.Client{
					Watcher: tc.mock(),
				},
			}
			ch, err := r.Subscribe("service-name")
			assert.Equal(t, tc.wantErr, err)
			event := <-ch
			// 你可以在这里进一步断言你预期中event, 还要测试close的例子
			log.Println(event)
		})
	}
}
