package c9_v1

import (
	"fmt"
	"testing"
)

func Test_Map(t *testing.T) {
	var m = make(map[mapKey]string)
	var key = mapKey{10}

	m[key] = "hello"
	fmt.Printf("m[key]=%s\n", m[key]) // ==> m[key]=hello

	// 修改key的字段的值后再次查询map, 无法获取刚才add进去的值
	key.key = 100
	fmt.Printf("m[key]=%s\n", m[key]) // ==> m[key]=
}

type mapKey struct {
	key int
}
