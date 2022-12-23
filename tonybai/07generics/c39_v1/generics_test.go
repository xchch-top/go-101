package c39_v1

import (
	"fmt"
	"testing"
)

func Test_Generics(t *testing.T) {
	fmt.Println(maxInt([]int{1, 2, -4, -6, 7, 0}))                       // 输出：7
	fmt.Println(maxString([]string{"11", "22", "44", "66", "77", "10"})) // 输出：77
	fmt.Println(maxFloat([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
}

func maxInt(sl []int) int {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxString(sl []string) string {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxFloat(sl []float64) float64 {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// go test -bench . generics_test.go
