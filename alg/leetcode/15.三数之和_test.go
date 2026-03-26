package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test_threeSum",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Test_threeSum",
			args: args{nums: []int{3, 0, -2, -1, 1, 2}},
			want: [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "Test_threeSum",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}},
			want: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, threeSum(tt.args.nums), "threeSum(%v)", tt.args.nums)
		})
	}
}
