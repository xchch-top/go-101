package main

import (
	rpc "gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v3"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v3/serialize/json"
	"gitlab.xchch.top/zhangsai/go-101/training/week14/rpc_v2/v3/serialize/proto"
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