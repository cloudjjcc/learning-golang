package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_pivotIndex",
			args: args{nums: []int{1, 7, 3, 6, 5, 6}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pivotIndex(tt.args.nums), "pivotIndex(%v)", tt.args.nums)
		})
	}
}
