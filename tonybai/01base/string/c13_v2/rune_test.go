package c13_v2

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// rune -> []byte
func Test_EncodeRune(t *testing.T) {
	var r rune = 0x4E2D
	fmt.Printf("the unicode charactor is %c\n", r) // 中
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)                       // 对rune进行utf-8编码
	fmt.Printf("utf-8 representation is 0x%X\n", buf) // 0xE4B8AD
}

// []byte -> rune
func Test_DecodeRune(t *testing.T) {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)                                                             // 对buf进行utf-8解码
	fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r)) // 中
}
