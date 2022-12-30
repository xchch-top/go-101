package render

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
	"testing"
)

var (
	msg = TestStruct{
		Name:    "Lena",
		Message: "hey",
		Number:  123,
	}
)

// 将输出渲染为json格式
func Test_Json_Render(t *testing.T) {
	render := gin.Default()

	render.GET("json-map", func(ctx *gin.Context) {
		// 使用map
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  "StatusOK",
		})
	})
	// 输出 {"message":"hey","status":"StatusOK"}

	render.GET("json-struct", func(ctx *gin.Context) {
		// 使用struct, 结构体中User通过json输出为user
		ctx.JSON(http.StatusOK, msg)
	})
	// 输出 {"user":"Lena","Message":"hey","Number":123}

	render.Run(":8080")
}

type TestStruct struct {
	Name    string `json:"user"`
	Message string
	Number  int
}

// 将输出渲染为xml格式
func Test_Xml_Render(t *testing.T) {
	render := gin.Default()

	render.GET("xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, msg)
	})
	// 输出 <TestStruct><Name>Lena</Name><Message>hey</Message><Number>123</Number></TestStruct>

	render.Run(":8080")
}

// 将输出渲染为yaml格式
func Test_Yaml_Render(t *testing.T) {
	render := gin.Default()

	render.GET("yaml", func(ctx *gin.Context) {
		ctx.YAML(http.StatusOK, msg)
	})
	// 输出
	// name: Lena
	// message: hey
	// number: 123

	render.Run(":8080")
}

// 将输出渲染为protobuf格式
func Test_Protobuf_Render(t *testing.T) {
	render := gin.Default()

	render.GET("protobuf", func(ctx *gin.Context) {
		resp := []int64{int64(1), int64(2)}
		label := "test"

		data := &protoexample.Test{
			Label: &label,
			Reps:  resp,
		}
		// 数据在响应中变为二进制数据
		ctx.ProtoBuf(http.StatusOK, data)
	})
	// 输出 test

	render.Run(":8080")
}
