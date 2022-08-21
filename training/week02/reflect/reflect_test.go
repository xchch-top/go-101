package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectPanic01(t *testing.T) {
	typ := reflect.TypeOf(123)
	// panic
	// typ.NumField()

	if typ.Kind() == reflect.Struct {
		// 不会执行到这里
		n := typ.NumField()
		fmt.Println(n)
	}
}

func TestReflectPanic02(t *testing.T) {
	typ := reflect.TypeOf(User{})

	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		fmt.Println(n)
	}
}

func TestReflectPanic03(t *testing.T) {
	typ := reflect.TypeOf(&User{})

	if typ.Kind() == reflect.Struct {
		// 不会执行到这里
		n := typ.NumField()
		fmt.Println(n)
	}
}

func TestReflectPanic04(t *testing.T) {
	typ := reflect.TypeOf(&User{})

	// panic
	n := typ.NumField()
	fmt.Println(n)
}

func TestReflectPanic05(t *testing.T) {
	typ := reflect.TypeOf(&User{})

	if typ.Kind() == reflect.Struct {
		fmt.Println("结构体")
	} else if typ.Kind() == reflect.Ptr {
		fmt.Println("指针")
	}
}

type User struct {
	name string
}
