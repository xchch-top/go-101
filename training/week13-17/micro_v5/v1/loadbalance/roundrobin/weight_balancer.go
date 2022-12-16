package roundrobin

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"sync"
)

type WeightPicker struct {
	sync.Mutex
	conns []*conn
}

func (w *WeightPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	if len(w.conns) == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	w.Lock()
	defer w.Unlock()
	// 接下来就是执行算法了
	var totalWeight uint32
	var maxWeightConn *conn
	for _, c := range w.conns {
		efficientWeight := c.efficientWeight
		totalWeight = totalWeight + efficientWeight
		c.currentWeight = c.currentWeight + efficientWeight
		if maxWeightConn == nil || maxWeightConn.currentWeight < c.currentWeight {
			// 这是挑选
			maxWeightConn = c
		}
	}
	maxWeightConn.currentWeight = maxWeightConn.currentWeight - totalWeight

	return balancer.PickResult{
		SubConn: maxWeightConn.SubConn,
		Done: func(info balancer.DoneInfo) {
			w.Lock()
			defer w.Unlock()

			// 如果要动态调整权重, 最好是设置一个上限和一个下限 [0, weight]
			// 考虑极端情况, 调整权重会不会导致某个节点永远选不上, 或者某个节点一直选上

			// 调整efficient weight
			// if info.Err != nil {
			// 	maxWeightConn.efficientWeight = maxWeightConn.efficientWeight - 1
			// } else {
			// 	maxWeightConn.efficientWeight = maxWeightConn.efficientWeight + 1
			// }
		},
	}, nil
}

type WeightBuilder struct {
}

func (w *WeightBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	conns := make([]*conn, 0, len(info.ReadySCs))
	for subConn, subConnInfo := range info.ReadySCs {
		weight := subConnInfo.Address.Attributes.Value("weight").(uint32)
		conns = append(conns, &conn{
			SubConn: subConn,
			// 你怎么得到权重?
			weight:          weight,
			currentWeight:   weight,
			efficientWeight: weight,
		})
	}

	return &WeightPicker{
		conns: conns,
	}
}

func (p *WeightBuilder) Name() string {
	return "WEIGHT_ROUND_ROBIN"
}

type conn struct {
	balancer.SubConn
	weight          uint32
	currentWeight   uint32
	efficientWeight uint32
}
