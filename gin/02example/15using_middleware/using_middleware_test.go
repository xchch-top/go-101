package using_middleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_Using_Middleware(t *testing.T) {
	// 新建一个没有任何默认中间件的路由
	router := gin.New()

	// 全局中间件
	// Logger中间件将日志写入 gin.DefaultWriter, 默认值gin.DefaultWriter = os.Writer
	router.Use(gin.Logger())
	// Recovery中间件会recover任何panic, 如果有panic的话, 会写入500
	router.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用下面两行代码效果一样
	authorized := router.Group("/")
	// 路由组中间件, 在此例中, 我们在authorized路由组中使用自定义创建的AuthRequired()中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	router.Run(":8080")
}

func analyticsEndpoint(ctx *gin.Context) {

}

func readEndpoint(ctx *gin.Context) {

}

func submitEndpoint(ctx *gin.Context) {

}

func loginEndpoint(ctx *gin.Context) {

}

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func benchEndpoint(ctx *gin.Context) {
}

func MyBenchLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
