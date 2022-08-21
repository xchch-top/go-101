package reflect

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateFields(t *testing.T) {
	s1 := &Stu{
		Name: "multiple pointer",
	}
	s2 := &s1

	tests := []struct {
		// 名字
		name string

		// 输入部分
		val any

		// 输出部分
		wantRes map[string]any
		wantErr error
	}{
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("不能为 nil"),
		},
		{
			name:    "stu",
			val:     Stu{Name: "TOM"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "TOM",
			},
		},
		{
			name: "pointer",
			val:  &Stu{Name: "Jerry"},
			// 要支持指针
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Jerry",
			},
		},
		{
			name: "multiple pointer",
			val:  s2,
			// 要支持指针
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "multiple pointer",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := iterateFields(tt.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestSetField(t *testing.T) {
	s1 := &Stu{
		Name: "zhangsan",
	}

	err := SetField(s1, "Name", "lisi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1.Name)
}
