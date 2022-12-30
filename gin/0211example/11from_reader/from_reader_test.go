package from_reader

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

// 从reader读取数据
func Test_From_Reader(t *testing.T) {
	router := gin.Default()

	router.GET("/from-reader", func(ctx *gin.Context) {
		response, err := http.Get("https://profile.csdnimg.cn/A/2/6/1_weixin_42503264")
		if err != nil || response.StatusCode != http.StatusOK {
			log.Println(err.Error())
			ctx.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.Run(":8080")
}
