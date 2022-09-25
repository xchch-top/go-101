package v3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.xchch.top/zhangsai/go-101/training/week07/orm_v2/v3/internal/errs"
	"reflect"
	"testing"
)

func Test_parseModel(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    *model
		wantErr error
	}{
		{
			name:  "ptr",
			input: &TestModel{},
			want: &model{
				tableName: "test_model",
				fieldMap: map[string]*field{
					"Id": {
						colName: "id",
					},
					"FirstName": {
						colName: "first_name",
					},
					"Age": {
						colName: "age",
					},
					"LastName": {
						colName: "last_name",
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
			want: &model{
				tableName: "test_model",
				fieldMap: map[string]*field{
					"Id": {
						colName: "id",
					},
					"FirstName": {
						colName: "first_name",
					},
					"Age": {
						colName: "age",
					},
					"LastName": {
						colName: "last_name",
					},
				},
			},
		},
	}

	r := &registry{
		models: map[reflect.Type]*model{},
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

	// hello, test model <nil>
	// hello, test model nil
	fSwtich((*TestModel)(nil))
}

func fSwtich(val any) {
	switch v := val.(type) {
	case nil:
		fmt.Println("hello, nil")
	case *TestModel:
		fmt.Printf("hello, test model %v \n", v)
		if v == nil {
			fmt.Println("hello, test model nil")
		}
	}
}
