package main

import (
	"context"
	micro "gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v1"
	"gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v1/example/proto/gen"
	"gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v1/registry"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var r registry.Registry
	rsBuilder := micro.NewResolverBuilder(r)
	cc, err := grpc.Dial("registry:///user-service",
		grpc.WithInsecure(), grpc.WithResolvers(rsBuilder))
	if err != nil {
		panic(err)
	}

	usClient := gen.NewUserServiceClient(cc)
	resp, err := usClient.GetById(context.Background(), &gen.GetByIdReq{
		Id: 12,
	})
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
