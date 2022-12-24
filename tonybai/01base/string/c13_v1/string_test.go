package c13_v1

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Test_String_V1(t *testing.T) {
	// 字节视角
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s)) // 9
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
}

func Test_String_V2(t *testing.T) {
	// 字符视角
	var s = "中国人"
	fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3
	for _, c := range s {
		fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
	}
}
