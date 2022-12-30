package static_file

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_static_file(t *testing.T) {
	router := gin.Default()
	// 用于静态资源访问
	router.Static("/assets", "./assets")
	// 用于文件下载的目录映射
	router.StaticFS("/download", http.Dir("D:\\gitlab-clone\\cs-basic-group\\halo-blog\\image"))
	// 请求站点图标
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.Run(":8080")
}
