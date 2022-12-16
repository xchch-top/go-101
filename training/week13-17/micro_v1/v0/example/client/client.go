package main

import (
	"context"
	micro "gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v0"
	gen2 "gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v0/example/proto/gen"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v0/registry"
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

	usClient := gen2.NewUserServiceClient(cc)
	resp, err := usClient.GetById(context.Background(), &gen2.GetByIdReq{
		Id: 12,
	})
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
