package v1

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
			wantErr: errors.New("orm: 只支持指针"),
		},
		{
			name:    "map",
			input:   map[string]string{},
			wantErr: errors.New("orm: 只支持指针"),
		},
		{
			name:    "nil",
			input:   nil,
			wantErr: errors.New("orm: 不支持nil"),
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseModel(tt.input)
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
