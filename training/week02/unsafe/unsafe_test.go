package unsafe

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnsafeAccessor_Field(t *testing.T) {

	testCases := []struct {
		name    string
		entity  interface{}
		field   string
		wantVal int
		wantErr error
	}{
		{
			name:    "invalid field",
			entity:  &User{age: 18},
			field:   "AgeXXX",
			wantErr: errors.New("不存在的字段"),
		},
		{
			name:    "normal case",
			entity:  &User{age: 18},
			field:   "age",
			wantVal: 18,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewUnsafeAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			val, err := accessor.Field(tc.field)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestUnsafeAccessor_SetField(t *testing.T) {
	type args struct {
		field string
		val   int
	}

	testCases := []struct {
		name    string
		entity  interface{}
		args    args
		wantVal int
		wantErr error
	}{
		{
			name:   "invalid field",
			entity: &User{age: 18},
			args: args{
				field: "AgeXXX",
				val:   20,
			},
			wantErr: errors.New("不存在的字段"),
		},
		{
			name:   "normal case",
			entity: &User{age: 18},
			args: args{
				field: "age",
				val:   20,
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewUnsafeAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			err = accessor.SetField(tc.args.field, tc.args.val)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
