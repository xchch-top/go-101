package main

import (
	"gitlab.xchch.top/golang-group/go-101/training/micro/rpc"
	"gitlab.xchch.top/golang-group/go-101/training/micro/rpc/serialize/json"
	"gitlab.xchch.top/golang-group/go-101/training/micro/rpc/serialize/proto"
)

func main() {
	svr := rpc.NewServer()
	svr.RegisterService(&UserService{})
	svr.RegisterService(&UserServiceProto{})
	svr.RegisterSerializer(json.Serializer{})
	svr.RegisterSerializer(proto.Serializer{})
	if err := svr.Start(":8081"); err != nil {
		panic(err)
	}
}
