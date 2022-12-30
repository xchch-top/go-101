package bind_uri

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_Bind_Uri(t *testing.T) {
	router := gin.Default()

	// uri ==> http://localhost:8080/zhang3/987fbc97-4bed-5078-9f07-9141ba07c9f3
	router.GET("/:name/:id", func(ctx *gin.Context) {
		var person Person
		err := ctx.ShouldBindUri(&person)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"name": person.Name, "id": person.Id})
	})

	router.Run(":8080")
}

type Person struct {
	Id   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}
