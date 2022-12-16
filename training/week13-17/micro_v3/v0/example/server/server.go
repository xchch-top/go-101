package main

import (
	"fmt"
	micro "gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v3/v0"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v3/v0/example/proto/gen"
	"gitlab.xchch.top/golang-group/go-101/training/week13-17/micro_v3/v0/registry/etcd"
	clientV3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/sync/errgroup"
	"strconv"
	"time"
)

func main() {
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
	var eg errgroup.Group

	for i := 0; i < 3; i++ {
		idx := i
		eg.Go(func() error {
			server := micro.NewServer("user-service", micro.ServerWithRegistry(r), micro.ServerWithTimeout(time.Second*3),
				micro.ServerWithWeight(uint32(1+idx)))
			defer server.Close()

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
	fmt.Println(err)
}
