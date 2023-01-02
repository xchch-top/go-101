package custom_validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testing"
	"time"
)

// 自定义验证器
func Test_Custom_Validators(t *testing.T) {
	router := gin.Default()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// url ==> http://localhost:8080/book?check_in=2023-01-01&check_out=2023-01-02
	router.GET("/book", func(ctx *gin.Context) {
		var b Booking
		err := ctx.ShouldBindWith(&b, binding.Query)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "date is valid"})
	})

	router.Run(":8080")
}

// Booking 包含绑定和验证数据
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		now := time.Now()
		if date.Before(now) {
			return false
		}
	}
	return true
}
