package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "", args: args{nums: []int{1, 1, 1, 2, 2, 3}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeDuplicates2(tt.args.nums), "removeDuplicates2(%v)", tt.args.nums)
		})
	}
}
