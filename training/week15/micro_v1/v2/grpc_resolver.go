package v2

import (
	"context"
	"gitlab.xchch.top/golang-group/go-101/training/week15/micro_v1/v2/registry"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

type grpcResolverBuilder struct {
	r       registry.Registry
	timeout time.Duration
}

func NewResolverBuilder(r registry.Registry, timeout time.Duration) resolver.Builder {
	return &grpcResolverBuilder{
		r: r, timeout: timeout,
	}
}

func (g *grpcResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {
	res := &grpcResolver{
		target:  target,
		cc:      cc,
		r:       g.r,
		close:   make(chan struct{}),
		timeout: g.timeout,
	}

	return res, res.watch()
}

func (g *grpcResolverBuilder) Scheme() string {
	return "registry"
}

type grpcResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	r      registry.Registry
	close  chan struct{}

	timeout time.Duration
}

// ResolveNow 立刻解析 == 立刻执行服务发现 == 立刻去问下注册中心
func (g *grpcResolver) ResolveNow(options resolver.ResolveNowOptions) {
	g.resolve()
}

func (g *grpcResolver) watch() error {
	eventsCh, err := g.r.Subscribe(g.target.Endpoint)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-eventsCh:
				// 做法一: 立即更新可用节点列表
				// 这种是幂等的
				log.Println(event)
				g.resolve()

				// 做法二: 精细化做法, 非常依赖于事件的顺序
				// 你这里收到的事件的顺序, 要和在注册中心上发生的顺序一样
				// 好处: 少访问一次注册中心
				// switch event.Type {
				// case registry.EventTypeAdd:
				// 	event.Instance // 这是新加节点
				// case registry.EventTypeDelete:
				// 	event.Instance // 这是被删除的节点
				// case registry.EventTypeUpdate:
				// 	event.Instance // 这是被更新的节点, 而且是更新后的节点
				// }
			case <-g.close:
				close(g.close)
				return
			}
		}
	}()
	return nil
}

func (g *grpcResolver) resolve() resolver.State {
	r := g.r
	// 这个就是可用服务实例(节点)列表
	// 考虑设置超时
	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	instances, err := r.ListService(ctx, g.target.Endpoint)
	if err != nil {
		g.cc.ReportError(err)
		return resolver.State{}
	}

	// 我是不是要把instances转成Address
	address := make([]resolver.Address, 0, len(instances))
	for _, ins := range instances {
		address = append(address, resolver.Address{
			// 定位信息 ip+端口
			Addr: ins.Address,
			// 可能还有其它字段
		})
	}
	state := resolver.State{Addresses: address}
	err = g.cc.UpdateState(state)
	if err != nil {
		g.cc.ReportError(err)
	}

	return state
}

func (g *grpcResolver) Close() {
	g.close <- struct{}{}
}
