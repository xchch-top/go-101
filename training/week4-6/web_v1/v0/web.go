package v0

import "net/http"

func Start() {
	http.ListenAndServe(":8081", &MyHandler{})
}

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}
