package v1

import (
	"log"
	"net"
	"net/http"
)

type HandleFunc func(*Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// AddRoute 注册路由的核心抽象
	AddRoute(method, path string, handler HandleFunc)
}

type HttpServer struct {
	router
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: newRouter(),
	}
}

func (m *HttpServer) AddRoute(method, path string, handler HandleFunc) {
	m.addRoute(method, path, handler)
}

func (m *HttpServer) Get(path string, handler HandleFunc) {
	m.AddRoute(http.MethodGet, path, handler)
}

func (m *HttpServer) Post(path string, handler HandleFunc) {
	m.AddRoute(http.MethodPost, path, handler)
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

func (m *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}

	// 接下来就是 查找路由、执行业务逻辑
	m.Serve(ctx)
}

func (m *HttpServer) Serve(ctx *Context) {
	mi, ok := m.router.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok {
		ctx.Resp.WriteHeader(http.StatusNotFound)
		ctx.Resp.Write([]byte("404 了"))
		return
	}

	ctx.PathParams = mi.pathParams
	mi.n.handler(ctx)
}
