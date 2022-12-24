package c20_v2

import "testing"

func Test_Switch(t *testing.T) {
	// x是一个接口类型变量, 它的静态类型为interface{}, 如果我们将整形值13赋值给x, x这个接口变量的动态类型就是int
	var x interface{} = 13
	switch x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int") // √
	case string:
		println("the type of x is string")
	case bool:
		println("the type of x is string")
	default:
		println("don't support the type")
	}
}

func Test_Switch_V2(t *testing.T) {
	var x interface{} = 13
	// x存储值信息
	switch v := x.(type) {
	case nil:
		println("v is nil")
	case int:
		println("the type of v is int, v =", v) // the type of v is int, v = 13
	case string:
		println("the type of v is string, v =", v)
	case bool:
		println("the type of v is bool, v =", v)
	default:
		println("don't support the type")
	}
}

func Test_Switch_V3(t *testing.T) {
	// var t T
	// var i I = t
	// switch i.(type) {
	// case T: // √
	// 	println("it is type T")
	// case int: // Impossible type switch case: 'int' does not implement 'I'
	// 	println("it is type int")
	// case string: // Impossible type switch case: 'string' does not implement 'I'
	// 	println("it is type string")
	// }
}

type I interface {
	M()
}

type T struct {
}

func (T) M() {
}
