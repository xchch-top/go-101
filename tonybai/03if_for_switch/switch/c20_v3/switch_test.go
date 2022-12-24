package c20_v3

import "testing"

func Test_Switch(t *testing.T) {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	// find first even number of the interger slice
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			// break语句只跳出了switch语句
			break
		case 1:
			// do nothing
		}
	}
	println(firstEven)
}
