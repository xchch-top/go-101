package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.xchch.top/golang-group/go-101/training/week15/micro_v1/v2/registry"
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
	client *clientV3.Client
	sess   *concurrency.Session

	mutex       sync.RWMutex
	watchCancel []func()
	close       chan struct{}
}

func NewRegistry(client *clientV3.Client) (*Registry, error) {
	sess, err := concurrency.NewSession(client)
	if err != nil {
		return nil, err
	}
	return &Registry{
		client: client,
		sess:   sess,
	}, nil
}

func (r *Registry) Register(ctx context.Context, ins registry.ServiceInstance) error {
	// ctx = clientV3.WithRequireLeader(ctx)
	// 准备 key value 和 租约
	instanceKey := fmt.Sprintf("/micro/%s/%s", ins.ServiceName, ins.Address)
	val, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	// TODO 手动管理租约, 要考虑续约间隔, 续约时长, 续约容错, 续约容错的过程对服务发现的影响
	_, err = r.client.Put(ctx, instanceKey, string(val), clientV3.WithLease(r.sess.Lease()))
	return err
}

func (r *Registry) Unregister(ctx context.Context, ins registry.ServiceInstance) error {
	instanceKey := fmt.Sprintf("/micro/%s/%s", ins.ServiceName, ins.Address)
	_, err := r.client.Delete(ctx, instanceKey)
	return err
}

func (r *Registry) ListService(ctx context.Context, serviceName string) ([]registry.ServiceInstance, error) {
	serviceKey := fmt.Sprintf("/micro/%s", serviceName)
	resp, err := r.client.Get(ctx, serviceKey, clientV3.WithPrefix())
	if err != nil {
		return nil, err
	}

	instances := make([]registry.ServiceInstance, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		// 你还不知道key是什么, value是什么
		var ins registry.ServiceInstance
		err = json.Unmarshal(kv.Value, &ins)
		if err != nil {
			// 你是跳过还是返回error呢?
			return nil, err
		}
		instances = append(instances, ins)
	}

	return instances, nil
}

func (r *Registry) Subscribe(serviceName string) (<-chan registry.Event, error) {
	serviceKey := fmt.Sprintf("/micro/%s", serviceName)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = clientV3.WithRequireLeader(ctx)
	r.mutex.Lock()
	r.watchCancel = append(r.watchCancel, cancel)
	r.mutex.Unlock()
	watchChan := r.client.Watch(ctx, serviceKey, clientV3.WithPrefix())
	res := make(chan registry.Event)

	go func() {
		for {
			select {
			case resp := <-watchChan:
				if resp.Canceled {
					return
				}
				if resp.Err() != nil {
					continue
				}
				for _, event := range resp.Events {
					// 在这里处理监听到的任何一个事件
					var ins registry.ServiceInstance
					err := json.Unmarshal(event.Kv.Value, &ins)
					if err != nil {
						// 忽略这个事件吗? 还是上报error, 怎么上报error呢?
						select {
						case res <- registry.Event{}:
						case <-ctx.Done():
							return
						}
						continue
					}
					select {
					case res <- registry.Event{
						Type:     typesMap[event.Type],
						Instance: ins,
					}:
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return res, nil
}

func (r *Registry) Close() error {
	r.mutex.RLock()
	watchCancel := r.watchCancel
	r.mutex.RUnlock()
	for _, cancel := range watchCancel {
		cancel()
	}

	_ = r.sess.Close()
	return nil
}
