package registry

import (
	"context"
)

type Registry interface {
	Register(ctx context.Context, ins ServiceInstance) error
	Unregister(ctx context.Context, ins ServiceInstance) error
	ListService(ctx context.Context, serviceName string) ([]ServiceInstance, error)
	Subscribe(serviceName string) (<-chan Event, error)
	// 可有可无, 不定义的话, 具体的实现也可以额外的添加
	Close() error
	// 可以考虑利用ctx来close掉返回的channel
	// Subscribe(ctx context.Context, serviceName string) (<-chan Event, error)
}

// ServiceInstance 代表的是一个实例
type ServiceInstance struct {
	Address string
}

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeAdd
	EventTypeDelete
	EventTypeUpdate
)

type Event struct {
	Type     EventType
	Instance ServiceInstance
}
