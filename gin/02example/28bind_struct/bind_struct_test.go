package bind_struct

import (
	"github.com/gin-gonic/gin"
	"testing"
)

// 绑定表单数据至自定义结构体
func Test_Bind_Struct(t *testing.T) {
	router := gin.Default()

	// url ==> http://localhost:8080/get-1?field_a=hello&field_b=world
	router.GET("/get-1", GetData1)

	// url ==> http://localhost:8080/get-2?field_a=hello&field_c=world
	router.GET("/get-2", GetData2)

	// url ==> http://localhost:8080/get-3?fielx_a=hello&field_d=world
	router.GET("/get-3", GetData3)

	router.Run(":8080")
}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetData1(ctx *gin.Context) {
	var b StructB
	ctx.Bind(&b)
	ctx.JSON(200, gin.H{"a": b.NestedStruct, "b": b.FieldB})
	// {"a":{"FieldA":"hello"},"b":"world"}
}

func GetData2(ctx *gin.Context) {
	var c StructC
	ctx.Bind(&c)
	ctx.JSON(200, gin.H{"a": c.NestedStructPointer, "c": c.FieldC})
	// {"a":{"FieldA":"hello"},"c":"world"}
}

func GetData3(ctx *gin.Context) {
	var d StructD
	ctx.Bind(&d)
	ctx.JSON(200, gin.H{"x": d.NestedAnonStruct, "d": d.FieldD})
	// {"d":"world","x":{"FieldX":"hello"}}
}
