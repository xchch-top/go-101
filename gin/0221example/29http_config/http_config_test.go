package http_config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

// 自定义http配置
func Test_Http_Config(t *testing.T) {
	router := gin.Default()

	http.ListenAndServe(":8080", router.Handler())
}

func Test_Http_Config_V2(t *testing.T) {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
