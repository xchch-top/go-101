package v2

import (
	"context"
	"gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v2/registry"
	"google.golang.org/grpc"
	"net"
	"time"
)

type ServerOption func(server *Server)

type Server struct {
	name string
	*grpc.Server
	r registry.Registry
}

func NewServer(name string, opts ...ServerOption) *Server {
	res := &Server{
		name:   name,
		Server: grpc.NewServer(),
	}
	for _, opt := range opts {
		opt(res)
	}

	return res
}

func ServerWithRegistry(r registry.Registry) ServerOption {
	return func(server *Server) {
		server.r = r
	}
}

func (s *Server) Start(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 一定是先启动端口再注册, 严格地来说, 是服务都启动了, 才注册
	if s.r != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		err = s.r.Register(ctx, registry.ServiceInstance{
			ServiceName: s.name,
		})
		cancel()
		if err != nil {
			return err
		}
	}
	return s.Serve(listen)
}
