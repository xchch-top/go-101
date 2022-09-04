package v3

import (
	"log"
	"net"
	"net/http"
)

func Start() {
	var s Server = &HttpServer{}
	s.Start(":8081")
}

type Server interface {
	http.Handler
	Start(addr string) error
}

type HttpServer struct {
}

type HttpsServer struct {
	Server
	CertFile string
	KeyFile  string
}

func (m *HttpServer) Start(addr string) error {
	// 端口启动前
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 端口启动后
	// 可以注册本服务器实例到你的管理平台, 比如说你注册到etcd, 然后你打开管理界面, 你就能看到这个实例
	log.Println("成功监听端口 8081")

	return http.Serve(listener, m)
	// 这个是阻塞的
	// return http.ListenAndServe(":8081", m)
}

func (m *HttpsServer) Start(addr string) error {
	return http.ListenAndServeTLS(addr, m.CertFile, m.KeyFile, m)
}

func (m *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello v3"))
}
