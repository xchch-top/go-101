package main

import (
	rpc "gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v3/v1"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v3/v1/serialize/json"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v3/v1/serialize/proto"
)

func main() {
	svr := rpc.NewServer()
	svr.MustRegister(&UserService{})
	svr.MustRegister(&UserServiceProto{})
	svr.RegisterSerializer(json.Serializer{})
	svr.RegisterSerializer(proto.Serializer{})
	if err := svr.Start(":8081"); err != nil {
		panic(err)
	}
}
