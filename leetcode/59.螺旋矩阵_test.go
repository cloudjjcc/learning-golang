package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test_generateMatrix",
			args: args{n: 3},
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, generateMatrix(tt.args.n), "generateMatrix(%v)", tt.args.n)
		})
	}
}
