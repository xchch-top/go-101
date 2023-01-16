package pprof

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_Pprof(t *testing.T) {
	router := gin.Default()
	pprof.Register(router)

	router.GET("/ping", func(c *gin.Context) {
		// H是map[string]interface{}的别名
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
