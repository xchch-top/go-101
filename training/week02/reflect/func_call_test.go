package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFuncs(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*FuncInfo
		wantErr error
	}{
		{
			name:    "nil",
			wantErr: errors.New("输入 nil"),
		},
		{
			name: "basic types",
			args: args{
				val: 123,
			},
			wantErr: errors.New("不支持的类型"),
		},
		{
			name: "type struct",
			args: args{
				val: Order{
					buyer:  18,
					seller: 100,
				},
			},
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(Order{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(0))},
					Result: []any{int64(18)},
				},
				// "getSeller": {
				// 	Name:   "getSeller",
				// 	In:     []reflect.Type{reflect.TypeOf(Order{})},
				// 	Out:    []reflect.Type{reflect.TypeOf(int64(0))},
				// 	Result: []any{int64(100)},
				// },
			},
		},
		{
			name: "type pointer",
			args: args{
				val: &OrderV2{
					buyer: 18,
				},
			},
			want: map[string]*FuncInfo{
				"GetBuyerV2": {
					Name:   "GetBuyerV2",
					In:     []reflect.Type{reflect.TypeOf(&OrderV2{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(0))},
					Result: []any{int64(18)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFuncs(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

type Order struct {
	buyer  int64
	seller int64
}

func (o Order) GetBuyer() int64 {
	return o.buyer
}

func (o Order) getSeller() int64 {
	return o.seller
}

type OrderV2 struct {
	buyer int64
}

func (o *OrderV2) GetBuyerV2() int64 {
	return o.buyer
}
