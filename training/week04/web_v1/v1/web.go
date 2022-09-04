package v1

import "net/http"

func Start() {
	var s Server = &HttpServer{}
	http.ListenAndServe(":8081", s)
	http.ListenAndServeTLS(":8481", "xxx", "xxx", s)
}

type Server interface {
	http.Handler
}

type HttpServer struct {
}

func (m *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello v1"))
}
