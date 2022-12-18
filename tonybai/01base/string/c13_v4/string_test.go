package c13_v4

import (
	"fmt"
	"testing"
)

// 下标操作
func Test_String_V1(t *testing.T) {
	var s = "中国人"
	fmt.Printf("0x%x\n", s[0]) // 0xe4
	// 字符“中” utf-8编码的第一个字节
}

// 字符迭代
func Test_String_V2(t *testing.T) {
	var s = "中国人"
	// 常规for迭代得到的是字符串内容的字节
	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
	}

	fmt.Println()
	// for range迭代得到的是字符串中unicode字符的码点值
	for i, v := range s {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}
}

// 字符串转换: Go 支持字符串与字节切片、字符串与 rune 切片的双向转换，并且这种转换无需调用任何函数，只需使用显式类型转换就可以了
func Test_String_V3(t *testing.T) {
	var s = "中国人"

	// string -> []rune
	rs := []rune(s)
	fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]

	// string -> []byte
	bs := []byte(s)
	fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba

	// []rune -> string
	s1 := string(rs)
	fmt.Println(s1) // 中国人

	// []byte -> string
	s2 := string(bs)
	fmt.Println(s2) // 中国人
}

// 字符串比较
func Test_String_V4(t *testing.T) {

}

// 字符串转换
func Test_String_V5(t *testing.T) {

}
