package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Abs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func Test_Max(t *testing.T) {
	got := Max(1, 2)
	if got != 2 {
		t.Errorf("Max(1, 2) = %f; want 2", got)
	}
}

// 多个输入的测试用例
func Test_Abs_2(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		{name: "t1", x: -0.3, want: 0.3},
		{name: "t2", x: 2, want: 2},
		{name: "t3", x: -3.1, want: 3.1},
		{name: "t4", x: 5, want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Abs(tc.x)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Benchmark_RandInt(b *testing.B) {
	// 重置性能测试时间计数
	// b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RandInt()
	}
}
