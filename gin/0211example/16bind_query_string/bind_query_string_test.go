package bind_query_string

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

// ShouldBindQuery函数只绑定url查询参数, 而忽略post数据
func Test_Bind_Query_String(t *testing.T) {
	router := gin.Default()

	router.Any("/bind", func(ctx *gin.Context) {
		var person Person
		if ctx.ShouldBindQuery(&person) == nil {
			log.Printf("bind, name: %s, address: %s\n", person.Name, person.Address)
		}
		ctx.String(http.StatusOK, "Success")
	})

	router.Run(":8080")
}

// http://localhost:8080/bind?name=n&address=a
// log输出: bind, name: n, address: a

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}
