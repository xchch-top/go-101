package bind_query_or_post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

// 绑定查询字符串或表单数据
func Test_Bind_Query_Or_Post(t *testing.T) {
	router := gin.Default()

	// localhost:8080/bind?name=boy&address=xyz&birthday=1992-03-15
	router.GET("/bind", func(ctx *gin.Context) {
		var p Person
		// 如果是GET请求, 只使用Form绑定引擎
		// 如果是POST请求, 先检查content-type是否为json或者xml, 然后再使用form
		err := ctx.ShouldBind(&p)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"name": p.Name, "address": p.Address, "birthday": p.Birthday})
	})

	router.Run(":8080")
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}
