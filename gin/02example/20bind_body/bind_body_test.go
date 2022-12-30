package bind_body

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"testing"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// 一般通过 c.Request.Body方法绑定数据, 但不能多次调用这个方法
func Test_Bind_Body(t *testing.T) {
	router := gin.Default()

	router.POST("/bind", func(ctx *gin.Context) {
		var objA formA
		var objB formB

		errA := ctx.ShouldBind(&objA)
		if errA == nil {
			ctx.String(http.StatusOK, fmt.Sprintf("the value from formA is %s", objA.Foo))
		}
		errB := ctx.ShouldBind(&objB)
		if errB == nil {
			ctx.String(http.StatusOK, fmt.Sprintf("the value from formB is %s", objB.Bar))
		} else {
			// 第二次解析body时会报错 EOF
			ctx.String(http.StatusInternalServerError, errB.Error())
		}
	})

	router.Run(":8080")
}

// 想多次绑定, 可以使用 c.ShouldBindBodyWith
// c.ShouldBindBodyWith 会在绑定之前将body存储到上下文中, 这会对性能造成轻微影响, 如果调用一次就能完成绑定的话, 就不要用这个方法
// 只有某些格式需要此功能, 如json, xml, msgPack, Protobuf
// 对于其他格式, 如query, form, FormPost, FormMultipart 可以多次调用c.ShouldBind()而不会造成任何性能损失
func Test_Bind_Body_V2(t *testing.T) {
	router := gin.Default()

	router.POST("/bind", func(ctx *gin.Context) {
		var objA formA
		var objB formB

		errA := ctx.ShouldBindBodyWith(&objA, binding.JSON)
		if errA == nil {
			ctx.String(http.StatusOK, fmt.Sprintf("the value from formA is %s", objA.Foo))
		}
		errB := ctx.ShouldBindBodyWith(&objB, binding.JSON)
		if errB == nil {
			ctx.String(http.StatusOK, fmt.Sprintf("the value from formB is %s", objB.Bar))
		} else {
			ctx.String(http.StatusInternalServerError, errB.Error())
		}
	})

	router.Run(":8080")
}
