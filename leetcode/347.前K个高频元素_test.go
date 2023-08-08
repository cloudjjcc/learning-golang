package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_topKFrequent",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Test_topKFrequent",
			args: args{
				nums: []int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0},
				k:    6,
			},
			want: []int{-3, -4, 0, 1, 4, 9},
		},
		{
			name: "Test_topKFrequent",
			args: args{
				nums: []int{5, 1, -1, -8, -7, 8, -5, 0, 1, 10, 8, 0, -4, 3, -1, -1, 4, -5, 4, -3, 0, 2, 2, 2, 4, -2, -4, 8, -7, -7, 2, -8, 0, -8, 10, 8, -8, -2, -9, 4, -7, 6, 6, -1, 4, 2, 8, -3, 5, -9, -3, 6, -8, -5, 5, 10, 2, -5, -1, -5, 1, -3, 7, 0, 8, -2, -3, -1, -5, 4, 7, -9, 0, 2, 10, 4, 4, -4, -1, -1, 6, -8, -9, -1, 9, -9, 3, 5, 1, 6, -1, -2, 4, 2, 4, -6, 4, 4, 5, -5},
				k:    7,
			},
			want: []int{4, -1, 2, -5, -8, 0, 8},
		},
		{
			name: "Test_topKFrequent",
			args: args{
				nums: []int{3, 0, 1, 0},
				k:    1,
			},
			want: []int{0},
		},
		{
			name: "Test_topKFrequent",
			args: args{
				nums: []int{4, 1, -1, 2, -1, 2, 3},
				k:    2,
			},
			want: []int{-1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, topKFrequent2(tt.args.nums, tt.args.k), "topKFrequent(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
