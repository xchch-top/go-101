package mock

import (
	orm "gitlab.xchch.top/golang-group/go-101/training/week7-10/orm_v11/v0"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := &MiddlewareBuilder{}
	_, err := orm.Open("sqlite3", "file:test.db?cache=shared&mode=memory",
		orm.DbWithMiddleware(builder.Build()))
	if err != nil {
		t.Fatal(err)
	}

	// 测试你的业务代码
	// DoBusiness(context.WithValue(context.Background(), mockKey{}, &Mock{
	//	Result: &User{},
	// }))

}
