package v0

import (
	"log"
	"net"
	"net/http"
)

func Start() {
	var s Server = &HttpServer{}
	s.Start(":8081")

	var h1 HandleFunc
	var h2 HandleFunc
	s.AddRoute(http.MethodGet, "/user", func(ctx *Context) {
		h1(ctx)
		h2(ctx)
	})

	s.AddRoute(http.MethodGet, "/user", nil)
}

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter
}

type HandleFunc func(*Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// AddRoute 注册路由的核心抽象
	AddRoute(method, path string, handler HandleFunc)
}

type HttpServer struct {
}

type HttpsServer struct {
	Server
	CertFile string
	KeyFile  string
}

func (m *HttpServer) AddRoute(method, path string, handler HandleFunc) {

}

func (m *HttpServer) Get(path string, handler HandleFunc) {
	m.AddRoute(http.MethodGet, path, handler)
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
}

func (m *HttpsServer) Start(addr string) error {
	return http.ListenAndServeTLS(addr, m.CertFile, m.KeyFile, m)
}

func (m *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:    request,
		Writer: writer,
	}

	// 接下来就是 查找路由、执行业务逻辑
	m.Serve(ctx)
}

func (m *HttpServer) Serve(ctx *Context) {
	ctx.Writer.Write([]byte("hello v4"))
}
