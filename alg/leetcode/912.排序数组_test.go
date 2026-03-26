package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_sortArray2",
			args: args{nums: []int{5, 2, 3, 1}},
			want: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sortArray2(tt.args.nums), "sortArray2(%v)", tt.args.nums)
		})
	}
}
