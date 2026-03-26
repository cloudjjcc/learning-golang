package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_wiggleSort",
			args: args{nums: []int{3, 5, 2, 1, 6, 4}},
			want: []int{3, 5, 1, 6, 2, 4},
		},
		{
			name: "Test_wiggleSort",
			args: args{nums: []int{6, 6, 5, 6, 3, 8}},
			want: []int{6, 6, 5, 6, 3, 8},
		},
		{
			name: "Test_wiggleSort",
			args: args{nums: []int{2, 1}},
			want: []int{1, 2},
		},
		{
			name: "Test_wiggleSort",
			args: args{nums: []int{3, 1, 2}},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			assert.Equal(t, tt.args.nums, tt.want)
		})
	}
}
