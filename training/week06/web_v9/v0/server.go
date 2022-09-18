package v0

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
	ms        []Middleware
	tplEngine TemplateEngine
}

type ServerOption func(server *HttpServer)

func NewHttpServer(opts ...ServerOption) *HttpServer {
	s := &HttpServer{
		router: newRouter(),
	}
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func ServerWithTemplateEngine(engine TemplateEngine) ServerOption {
	return func(server *HttpServer) {
		server.tplEngine = engine
	}
}

func (s *HttpServer) Use(ms ...Middleware) {
	s.ms = ms
}

func (s *HttpServer) AddRoute(method, path string, handler HandleFunc) {
	s.addRoute(method, path, handler)
}

func (s *HttpServer) Get(path string, handler HandleFunc) {
	s.AddRoute(http.MethodGet, path, handler)
}

func (s *HttpServer) Post(path string, handler HandleFunc) {
	s.AddRoute(http.MethodPost, path, handler)
}

func (s *HttpServer) Start(addr string) error {
	// 端口启动前
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 端口启动后
	// 可以注册本服务器实例到你的管理平台, 比如说你注册到etcd, 然后你打开管理界面, 你就能看到这个实例
	log.Println("成功监听端口 8081")

	return http.Serve(listener, s)
}

func (s *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:       request,
		Resp:      writer,
		tplEngine: s.tplEngine,
	}

	root := s.serve
	for i := len(s.ms) - 1; i >= 0; i-- {
		m := s.ms[i]
		root = m(root)
	}

	// 将sMdl加为middleware的最后一个
	var flashMdl Middleware = func(next HandleFunc) HandleFunc {
		return func(ctx *Context) {
			next(ctx)
			s.writeResp(ctx)
		}
	}
	root = flashMdl(root)
	root(ctx)
}

func (s *HttpServer) writeResp(ctx *Context) {
	// ctx.Resp.Header().Set("Content-Type", "application/json")
	//ctx.Resp.WriteHeader(ctx.StatusCode)
	ctx.Resp.Write(ctx.RespData)
}

func (s *HttpServer) serve(ctx *Context) {
	mi, ok := s.router.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok {
		ctx.StatusCode = http.StatusNotFound
		ctx.RespData = []byte("404 了")
		return
	}
	ctx.PathParams = mi.pathParams
	ctx.MatchedRoute = mi.n.route
	mi.n.handler(ctx)
}
