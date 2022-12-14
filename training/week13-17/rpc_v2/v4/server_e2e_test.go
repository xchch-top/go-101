//go:build e2e

package v4

import (
	gen "gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v4/example/rpc/server"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v4/serialize/json"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/rpc_v2/v4/serialize/proto"
	"testing"
)

func TestServer_Start(t *testing.T) {
	s := NewServer()
	s.MustRegister(&gen.UserServiceProto{})
	s.RegisterSerializer(json.Serializer{})
	s.RegisterSerializer(proto.Serializer{})
	s.Start(":8081")
}
