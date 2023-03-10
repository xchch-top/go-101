package main

import (
	"context"
	micro "gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v2"
	gen2 "gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v2/example/proto/gen"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v1/v2/registry"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var r registry.Registry
	rsBuilder := micro.NewResolverBuilder(r, time.Second)
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
