package v1

import (
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v4/v1/registry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	name     string
	listener net.Listener

	si       registry.ServiceInstance
	registry registry.Registry
	// 单个操作的超时时间，一般用于和注册中心打交道
	timeout time.Duration
	*grpc.Server

	weight uint32
}

func ServerWithWeight(weight uint32) ServerOption {
	return func(server *Server) {
		server.weight = weight
	}
}

type ServerOption func(server *Server)

func NewServer(name string, opts ...ServerOption) *Server {
	res := &Server{
		name:   name,
		Server: grpc.NewServer(),
		// 最好就不要用0，有些算法可能在 0 的时候崩溃
		weight: 1,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func ServerWithRegistry(r registry.Registry) ServerOption {
	return func(server *Server) {
		server.registry = r
	}
}

func ServerWithTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func (s *Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = listener
	// 用户决定使用注册中心
	if s.registry != nil {
		s.si = registry.ServiceInstance{
			Name:    s.name,
			Address: listener.Addr().String(),
			Weight:  s.weight,
		}
		ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
		// 要确保端口启动之后才能注册
		err = s.registry.Register(ctx, s.si)
		cancel()
		if err != nil {
			return nil
		}
		// defer func() {
		// 	s.registry.UnRegister(ctx, s.si)
		// }()
	}
	return s.Server.Serve(listener)
}

func (s *Server) Close() error {
	// 这里一定是先从注册中心摘掉
	// 如果要是没有摘掉，那么就不能关闭后面的 listener
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	if s.registry != nil {
		if err := s.registry.UnRegister(ctx, s.si); err != nil {
			return err
		}
	}
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
