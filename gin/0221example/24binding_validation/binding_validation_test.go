package binding_validation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"testing"
)

// gin目前支持json, xml, yaml和标准表单值的绑定
// gin使用 https://github.com/go-playground/validator 进行验证
// 使用时, 需要在要绑定的素有的字段上, 设置相应的tag
// gin提供了两类绑定方法: Must Bind 和 Should Bind
//		Must Bind: 如果发生了错误, 则请求终止, 并触发c.AbortWithError(400, err).SetType(ErrorTypeBind), 如果您希望更好地控制绑定, 考虑使用ShouldBind等效方法
// 		Should Bind: 如果发生了错误, gin会返回错误并由开发者处理错误和请求
// 使用bind方法时, gin会尝试根据content-type判断如何绑定, 如果你明确知道要绑定什么, 可以使用MustBindWith或shouldBindWith
// 如果一个字段的tag加上了 bind:"required", 但是绑定时是控制, gin会报错

func Test_Binding_Validation(t *testing.T) {
	router := gin.Default()

	// 绑定json ==> {"user":"u","password":"p"}
	router.POST("/login", func(ctx *gin.Context) {
		var json Login
		err := ctx.ShouldBindJSON(&json)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if json.User != "u" || json.Password != "p" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "you are logged in"})
	})

	// 绑定xml ==> <?xml version="1.0" encoding="UTF-8"?><root><user>u</user><password>p</password></root>
	router.POST("/xml", func(ctx *gin.Context) {
		var xml Login
		err := ctx.ShouldBindXML(&xml)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if xml.User != "u" || xml.Password != "p" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "you are logged in"})
	})

	// 绑定form ==> Content-Type为application/x-www-form-urlencoded
	router.POST("/form", func(ctx *gin.Context) {
		var form Login
		err := ctx.ShouldBindWith(&form, binding.FormPost)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if form.User != "u" || form.Password != "p" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "you are logged in"})
	})

	router.Run(":8080")
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
