package main

import "fmt"

func main() {
	fmt.Println("遍历slice.................")
	test01()

	fmt.Println("遍历map.................")
	test02()

	fmt.Println("遍历字符串.................")
	test03()

	fmt.Println("遍历字符串.................")
	test04()

}

// 遍历slice
func test01() {
	sli := []int{1, 3, 5, 7, 9}
	for i, v := range sli {
		fmt.Printf("sli[%d] = %d \n", i, v)
	}
}

// 遍历map
func test02() {
	m := map[string]int{
		"bob":   1,
		"alice": 0,
	}
	for k, v := range m {
		fmt.Printf("k = %s, v = %d \n", k, v)
	}
}

// 遍历字符串
func test03() {

	title := "go语言第一课"
	for i, v := range title {
		fmt.Printf("%d ==> %s, 0x%x \n", i, string(v), v)
	}

}

// 遍历字符串
func test04() {

	title := "go语言第一课"
	for i := 0; i < len(title); i++ {
		fmt.Printf("%d ==> 0x%x \n", i, title[i])
	}

}
