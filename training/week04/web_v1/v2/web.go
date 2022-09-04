package v2

import "net/http"

func Start() {
	// var s Server = &HttpsServer{
	// 	Server: &HttpServer{},
	// }
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
	return http.ListenAndServe(addr, m)
}

func (m *HttpsServer) Start(addr string) error {
	return http.ListenAndServeTLS(addr, m.CertFile, m.KeyFile, m)
}

func (m *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello v2"))
}
