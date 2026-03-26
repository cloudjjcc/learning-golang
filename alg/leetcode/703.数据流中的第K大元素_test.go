package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest_Add(t *testing.T) {
	type args struct {
		k    int
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestKthLargest_Add",
			args: args{
				k:    3,
				vals: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 1, 1, 2, 3},
		},
		{
			name: "TestKthLargest_Add",
			args: args{
				k:    2,
				vals: []int{0, -1, 1, -2, -4, 3},
			},
			want: []int{1, 1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &KthLargest{h: &Heap{data: make([]int, 0, tt.args.k)}, k: tt.args.k}
			for i, val := range tt.args.vals {
				assert.Equal(t, tt.want[i], this.Add(val))
			}
		})
	}
}
