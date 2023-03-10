package roundrobin

import (
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v3/v2/loadbalance"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
	"sync/atomic"
)

type Picker struct {
	ins    []instance
	cnt    uint64
	filter loadbalance.Filter
}

func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	if len(p.ins) == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	// 做一个筛选
	candidates := make([]instance, 0, len(p.ins))
	for _, sub := range p.ins {
		if !p.filter(info, sub.address) {
			continue
		}
		candidates = append(candidates, sub)
	}

	cnt := atomic.AddUint64(&p.cnt, 1)
	index := cnt % uint64(len(candidates))
	return balancer.PickResult{
		SubConn: candidates[index].sub,
		Done: func(info balancer.DoneInfo) {
			// 可以根据调用结果来调整你的负载均衡策略
			// 假如说出错了
			if info.Err != nil {
				// 尝试把这个SubConn置为不可用, 或者临时移除出去
			}
		}}, nil
}

type PickerBuilder struct {
	Filter loadbalance.Filter
}

func (p *PickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	conns := make([]instance, 0, len(info.ReadySCs))
	for sub, subInfo := range info.ReadySCs {
		conns = append(conns, instance{
			sub:     sub,
			address: subInfo.Address,
		})
	}
	return &Picker{
		filter: p.Filter,
		ins:    conns,
	}
}

func (p *PickerBuilder) Name() string {
	return "ROUND_ROBIN"
}

type instance struct {
	sub     balancer.SubConn
	address resolver.Address
}
