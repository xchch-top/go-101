package main

import (
	"context"
	"fmt"
	micro "gitlab.xchch.top/golang-group/go-101/training/week15/micro_v2/v2"
	"gitlab.xchch.top/golang-group/go-101/training/week15/micro_v2/v2/example/proto/gen"
	"gitlab.xchch.top/golang-group/go-101/training/week15/micro_v2/v2/loadbalance/roundrobin"
	"gitlab.xchch.top/golang-group/go-101/training/week15/micro_v2/v2/registry/etcd"
	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"log"
	"time"
)

func main() {
	// 注册中心
	etcdClient, err := clientV3.New(clientV3.Config{
		Endpoints: []string{"192.168.1.10:12379"},
	})
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewRegistry(etcdClient)
	if err != nil {
		panic(err)
	}

	pickerBuilder := &roundrobin.WeightBuilder{}
	builder := base.NewBalancerBuilder(pickerBuilder.Name(), pickerBuilder, base.Config{HealthCheck: true})
	// 注册你的负载均衡策略
	balancer.Register(builder)
	cc, err := grpc.Dial("registry:///user-service",
		grpc.WithInsecure(),
		grpc.WithResolvers(micro.NewResolverBuilder(r, time.Second*3)),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, pickerBuilder.Name())))
	if err != nil {
		panic(err)
	}

	client := gen.NewUserServiceClient(cc)
	for i := 0; i < 100; i++ {
		resp, err := client.GetById(context.Background(), &gen.GetByIdReq{})
		if err != nil {
			panic(err)
		}
		log.Println(resp)
	}
}
