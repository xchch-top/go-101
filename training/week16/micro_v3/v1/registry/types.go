package registry

import (
	"context"
	"io"
)

type Registry interface {
	Register(ctx context.Context, si ServiceInstance) error
	UnRegister(ctx context.Context, si ServiceInstance) error

	ListServices(ctx context.Context, name string) ([]ServiceInstance, error)
	Subscribe(name string) <-chan Event

	io.Closer
}

type ServiceInstance struct {
	Name    string
	Address string
	Weight  uint32
	Group   string
}

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeAdd
	EventTypeDelete
)

type Event struct {
	Type     EventType
	Instance ServiceInstance
}
