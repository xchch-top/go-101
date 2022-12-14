package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.xchch.top/golang-group/go-101/training/week17/micro_v5/v0/registry"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"sync"
)

var typesMap = map[mvccpb.Event_EventType]registry.EventType{
	mvccpb.PUT:    registry.EventTypeAdd,
	mvccpb.DELETE: registry.EventTypeDelete,
}

type Registry struct {
	sess   *concurrency.Session
	client *clientV3.Client

	mutex       sync.RWMutex
	watchCancel []func()
}

func NewRegistry(c *clientV3.Client) (*Registry, error) {
	// 没有设置 ttl，所以默认是 60 秒，这个可以做成可配置的
	sess, err := concurrency.NewSession(c)
	if err != nil {
		return nil, err
	}
	return &Registry{
		sess:   sess,
		client: c,
	}, nil
}

func (r *Registry) Register(ctx context.Context, si registry.ServiceInstance) error {
	val, err := json.Marshal(si)
	if err != nil {
		return nil
	}
	// 这个 key 也可以做成可配置的
	_, err = r.client.Put(ctx, r.instanceKey(si),
		string(val), clientV3.WithLease(r.sess.Lease()))
	return err
}

func (r *Registry) UnRegister(ctx context.Context, si registry.ServiceInstance) error {
	_, err := r.client.Delete(ctx, r.instanceKey(si))
	return err
}

func (r *Registry) ListServices(ctx context.Context, name string) ([]registry.ServiceInstance, error) {
	resp, err := r.client.Get(ctx, r.serviceKey(name), clientV3.WithPrefix())
	if err != nil {
		return nil, err
	}
	res := make([]registry.ServiceInstance, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		var si registry.ServiceInstance
		err = json.Unmarshal(kv.Value, &si)
		if err != nil {
			return nil, err
		}
		res = append(res, si)
	}
	return res, nil
}

func (r *Registry) instanceKey(s registry.ServiceInstance) string {
	return fmt.Sprintf("/micro/%s/%s", s.Name, s.Address)
}

func (r *Registry) serviceKey(name string) string {
	return fmt.Sprintf("/micro/%s", name)
}

func (r *Registry) Subscribe(name string) <-chan registry.Event {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = clientV3.WithRequireLeader(ctx)
	r.mutex.Lock()
	r.watchCancel = append(r.watchCancel, cancel)
	r.mutex.Unlock()
	ch := r.client.Watch(ctx, r.serviceKey(name), clientV3.WithPrefix())
	// 这里有没有 Buffer 都无所谓
	res := make(chan registry.Event)
	// 监听变更
	go func() {
		for {
			select {
			case resp := <-ch:
				if resp.Canceled {
					return
				}
				if resp.Err() != nil {
					continue
				}
				for _, event := range resp.Events {
					res <- registry.Event{
						Type: typesMap[event.Type],
					}
				}
			case <-ctx.Done():
				return
			}
		}

	}()
	return res
}

func (r *Registry) Close() error {
	r.mutex.Lock()
	for _, cancel := range r.watchCancel {
		cancel()
	}
	r.mutex.Unlock()
	// 因为 client 是外面传进来的，所以我们这里不能关掉它。它可能被其它的人使用着
	return r.sess.Close()
}
