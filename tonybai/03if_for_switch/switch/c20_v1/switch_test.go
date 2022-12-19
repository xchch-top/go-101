package c20_v1

import "testing"

func Test_Switch(t *testing.T) {
	switch switchExpr() {
	case case1():
		println("exec case1")
	case case2_1(), case2_2():
		println("exec case2")
	case case3():
		println("exec case3")
	default:
		println("exec default")
	}
}

// 程序输出:
// eval switch expr
// eval case1 expr
// eval case2_1 expr
// eval case2_2 expr
// exec case2

// 1. go先对switch expr表达式进行求值
// 2. 然后再按case语句的出现顺序, 从上到校进行逐一求值, 在带有表达式列表的case语句中, go会从左到右, 对列表中的表达式进行求值
// 3. 如果switch表达式匹配到了某个case表达式, 那么程序就会执行这个case对应的代码分支, 这个分支后面的case表达式将不会再得到求值机会
// 4. 即便后面的case表达式的求值后也能与switch表达式匹配上, go也不会继续对这些表达式进行求值了

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}
func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchExpr() int {
	println("eval switch expr")
	return 2
}
