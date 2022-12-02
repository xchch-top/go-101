package main

import (
	"fmt"
	micro "gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v2"
	"gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v2/example/proto/gen"
	"gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v2/registry/etcd"
	clientV3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/sync/errgroup"
	"strconv"
)

func main() {
	etcdClient, err := clientV3.New(clientV3.Config{
		Endpoints: []string{"192.168.1.10:2379"},
	})
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewRegistry(etcdClient)
	if err != nil {
		panic(err)
	}

	var eg errgroup.Group
	for i := 0; i < 3; i++ {
		idx := i
		eg.Go(func() error {
			server := micro.NewServer("user-service", micro.ServerWithRegistry(r))
			us := &UserService{
				name: fmt.Sprintf("server-%d", idx),
			}
			gen.RegisterUserServiceServer(server, us)
			fmt.Println("启动服务器: " + us.name)
			// 端口分别是 8081, 8082, 8083
			return server.Start(":" + strconv.Itoa(8081+idx))
		})
	}
	// 正常或者异常退出都会返回
	err = eg.Wait()
}
