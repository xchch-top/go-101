package group_routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

// 路由组
func Test_Group_Routes(t *testing.T) {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("login", funGroup)
		v1.GET("register", funGroup)
	}

	v2 := router.Group("v2")
	{
		v2.GET("login", funGroup)
		v2.GET("register", funGroup)
	}

	router.Run(":8080")
}

func funGroup(ctx *gin.Context) {
	log.Printf("uri: %s", ctx.Request.RequestURI)
	ctx.String(http.StatusOK, "success")
}
