package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPrintFieldOffset(t *testing.T) {
	fmt.Println("~~~ User", unsafe.Sizeof(User{}))
	PrintFieldOffset(User{})

	fmt.Println("~~~ UserV1", unsafe.Sizeof(UserV1{}))
	PrintFieldOffset(UserV1{})

	fmt.Println("~~~ UserV2", unsafe.Sizeof(UserV2{}))
	PrintFieldOffset(UserV2{})
}
