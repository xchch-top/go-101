package v0

import (
	"context"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v0/registry"
	"google.golang.org/grpc/resolver"
	"log"
)

type grpcResolverBuilder struct {
	r registry.Registry
}

func NewResolverBuilder(r registry.Registry) resolver.Builder {
	return &grpcResolverBuilder{r: r}
}

func (g *grpcResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {
	eventsCh, err := g.r.Subscribe(target.Endpoint)
	if err != nil {
		return nil, err
	}

	res := &grpcResolver{
		target: target,
		cc:     cc,
		r:      g.r,
		close:  make(chan struct{}),
	}

	go func() {
		for {
			select {
			case event := <-eventsCh:
				// 做法一: 立即更新可用节点列表
				// 这种是幂等的
				log.Println(event)
				res.resolve()

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
			case <-res.close:

			}
		}
	}()

	return res, nil
}

func (g *grpcResolverBuilder) Scheme() string {
	return "registry"
}

type grpcResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	r      registry.Registry
	close  chan struct{}
}

// ResolveNow 立刻解析 == 立刻执行服务发现 == 立刻去问下注册中心
func (g *grpcResolver) ResolveNow(options resolver.ResolveNowOptions) {
	g.resolve()
}

func (g *grpcResolver) resolve() resolver.State {
	r := g.r
	// 这个就是可用服务实例(节点)列表
	instances, err := r.ListService(context.Background(), g.target.Endpoint)
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
