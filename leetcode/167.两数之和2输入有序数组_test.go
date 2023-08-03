package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum22(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_twoSum22",
			args: args{
				numbers: []int{-1, -1, 1, 1, 1, 1, 1, 1, 1},
				target:  -2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, twoSum22(tt.args.numbers, tt.args.target), "twoSum22(%v, %v)", tt.args.numbers, tt.args.target)
		})
	}
}
