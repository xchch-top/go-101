package c39_v1

import (
	"fmt"
	"testing"
)

func Test_Generic_V2(t *testing.T) {
	fmt.Println(maxAny([]any{1, 2, -4, -6, 7, 0}))                 // 输出：7
	fmt.Println(maxAny([]any{"11", "22", "44", "66", "77", "10"})) // 输出：77
	fmt.Println(maxAny([]any{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
}

func maxAny(sl []any) any {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		switch v.(type) {
		case int:
			if v.(int) > max.(int) {
				max = v
			}
		case string:
			if v.(string) > max.(string) {
				max = v
			}
		case float64:
			if v.(float64) > max.(float64) {
				max = v
			}
		}
	}
	return max
}
