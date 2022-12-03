package c4_v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Slice_Queue(t *testing.T) {
	testCases := []struct {
		name string
		in   []uint64
		out  []uint64
	}{
		{
			name: "success",
			in:   []uint64{22, 37},
			out:  []uint64{22, 37},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewSliceQueue(len(tc.in))
			for _, i := range tc.in {
				q.Enqueue(i)
			}

			out := make([]uint64, 0, len(tc.out))
			for range tc.out {
				out = append(out, q.Dequeue().(uint64))
			}
			assert.Equal(t, out, tc.out)
		})
	}
}
