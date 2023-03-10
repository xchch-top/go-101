package v2

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v2/registry"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v2/registry/mocks"
	"google.golang.org/grpc/resolver"
	"testing"
	"time"
)

func Test_grpcResolverBuilder_Build(t *testing.T) {
	testCases := []struct {
		name string
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func Test_grpcResolver_ResolverNow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name string

		mock func() registry.Registry

		wantState resolver.State
		wantErr   error
	}{
		{
			name: "resolver success",
			mock: func() registry.Registry {
				r := mocks.NewMockRegistry(ctrl)
				r.EXPECT().ListService(gomock.Any(), gomock.Any()).
					Return([]registry.ServiceInstance{
						{
							Address: "test-1",
						},
					}, nil)
				return r
			},
			wantState: resolver.State{
				Addresses: []resolver.Address{
					{
						Addr: "test-1",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cc := &mockClientConn{}
			rs := &grpcResolver{
				target: resolver.Target{},
				cc:     cc,
				r:      tc.mock(),
			}
			rs.ResolveNow(resolver.ResolveNowOptions{})
			state := cc.state
			assert.Equal(t, tc.wantErr, cc.err)
			if cc.err != nil {
				return
			}
			assert.Equal(t, tc.wantState, state)
		})
	}
}

type mockClientConn struct {
	state resolver.State
	err   error
	resolver.ClientConn
}

func (cc *mockClientConn) UpdateState(state resolver.State) error {
	cc.state = state
	return nil
}

func (cc *mockClientConn) ReportError(err error) {
	cc.err = err
}

func Test_grpcResolver_watch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name      string
		mock      func() (registry.Registry, chan registry.Event)
		wantErr   error
		wantState resolver.State
	}{
		{
			name: "wanted and close",
			mock: func() (registry.Registry, chan registry.Event) {
				r := mocks.NewMockRegistry(ctrl)
				ch := make(chan registry.Event)
				r.EXPECT().Subscribe(gomock.Any()).Return(ch, nil)
				r.EXPECT().ListService(gomock.Any(), gomock.Any()).
					Return([]registry.ServiceInstance{
						{
							Address: "test-1",
						},
					}, nil)
				return r, ch
			},
			wantState: resolver.State{
				Addresses: []resolver.Address{
					{
						Addr: "test-1",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		cc := &mockClientConn{}
		t.Run(tc.name, func(t *testing.T) {
			r, ch := tc.mock()
			closeCh := make(chan struct{})
			rs := &grpcResolver{
				r:     r,
				cc:    cc,
				close: closeCh,
			}
			err := rs.watch()
			assert.Equal(t, tc.wantErr, err)
			ch <- registry.Event{}

			time.Sleep(time.Second)
			// ??????????????????
			rs.Close()
			// ????????????????????? closeCh ????????? close ???
			_, ok := <-closeCh
			assert.False(t, ok)
		})
	}
}
