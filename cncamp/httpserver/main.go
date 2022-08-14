package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gitlab.xchch.top/zhangsai/go-101/cncamp/httpserver/xxutil"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 第四题 healthz 时，应返回200
	io.WriteString(w, "200")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 第一题: 接收客户端request, 并将request中带的header写入response header
	//   -- 假设自定义的header中带前缀 goh-, 请求工具用的postman
	for k, v := range r.Header {
		if strings.Contains(strings.ToLower(k), "goh-") {
			w.Header().Set(k, fmt.Sprintf("%s", v))
		}
	}

	// 第二题: 读取当前系统的环境变量中VERSION配置, 并写入response header
	s, b := os.LookupEnv("VERSION")
	if !b {
		fmt.Println("环境变量中 VERSION没有设置")
	}
	w.Header().Set("VERSION", s)

	// 第三题: Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	// 客户端返回状态码
	code := 200
	clientIp := xxutil.GetClientIP(r)
	fmt.Printf("clientIp: %s, code: %d\n", clientIp, code)

	io.WriteString(w, `{"code": code, "msg": "ok"}`)
}
