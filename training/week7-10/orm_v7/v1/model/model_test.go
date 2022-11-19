package model

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v7/v1/internal/errs"
	"reflect"
	"testing"
)

func Test_parseModel(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    *Model
		wantErr error
	}{
		{
			name:  "ptr",
			input: &TestModel{},
			want: &Model{
				TableName: "test_model",
				FieldMap: map[string]*Field{
					"Id": {
						ColName: "id",
					},
					"FirstName": {
						ColName: "first_name",
					},
					"Age": {
						ColName: "age",
					},
					"LastName": {
						ColName: "last_name",
					},
				},
			},
		},
		{
			name:    "struct",
			input:   TestModel{},
			wantErr: errs.ErrPointOnly,
		},
		{
			name:    "map",
			input:   map[string]string{},
			wantErr: errs.ErrPointOnly,
		},
		{
			name:    "nil",
			input:   nil,
			wantErr: errs.ErrInputNil,
		},
		{
			name:  "nil with type  -- 不过多研究",
			input: (*TestModel)(nil),
			want: &Model{
				TableName: "test_model",
				FieldMap: map[string]*Field{
					"Id": {
						ColName: "id",
					},
					"FirstName": {
						ColName: "first_name",
					},
					"Age": {
						ColName: "age",
					},
					"LastName": {
						ColName: "last_name",
					},
				},
			},
		},
		{
			name: "column tag",
			input: func() any {
				// 我们把测试结构体定义在方法内部, 防止被其他用例访问
				type ColumnTag struct {
					ID uint64 `orm:"column=id"`
				}
				return &ColumnTag{}
			}(),
			want: &Model{
				TableName: "column_tag",
				FieldMap: map[string]*Field{
					"ID": {
						ColName: "id",
					},
				},
			},
		},
		{
			name:  "table name ",
			input: &CustomTableName{},
			want: &Model{
				TableName: "custom_table_name_t",
				FieldMap: map[string]*Field{
					"Name": {
						ColName: "name",
					},
				},
			},
		},
		{
			name:  "table name ptr",
			input: &CustomTableNamePtr{},
			want: &Model{
				TableName: "custom_table_name_ptr_t",
				FieldMap: map[string]*Field{
					"Name": {
						ColName: "name",
					},
				},
			},
		},
	}

	r := &Registry{
		Models: map[reflect.Type]*Model{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.parseModel(tt.input)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

// 如果一个变量有类型, 它就不是nil, 但是在switch里把它转为类型之后, 再去判断值是否为nil, 就可能是nil
func Test_Nil(t *testing.T) {
	// hello, nil
	fSwtich(nil)

	// hello, test Model <nil>
	// hello, test Model nil
	fSwtich((*TestModel)(nil))
}

func fSwtich(val any) {
	switch v := val.(type) {
	case nil:
		fmt.Println("hello, nil")
	case *TestModel:
		fmt.Printf("hello, test Model %v \n", v)
		if v == nil {
			fmt.Println("hello, test Model nil")
		}
	}
}

type CustomTableName struct {
	Name string
}

func (c CustomTableName) TableName() string {
	return "custom_table_name_t"
}

type CustomTableNamePtr struct {
	Name string
}

func (c *CustomTableNamePtr) TableName() string {
	return "custom_table_name_ptr_t"
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
