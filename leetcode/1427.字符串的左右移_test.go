package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stringShift(t *testing.T) {
	type args struct {
		s     string
		shift [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_stringShift",
			args: args{s: "abc", shift: [][]int{{0, 1}, {1, 2}}},
			want: "cab",
		},
		{
			name: "Test_stringShift",
			args: args{s: "abcdefg", shift: [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}},
			want: "efgabcd",
		},
		{
			name: "Test_stringShift",
			args: args{s: "a", shift: [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}},
			want: "a",
		},
		{
			name: "Test_stringShift",
			args: args{s: "abc", shift: [][]int{{0, 1}, {1, 2}}},
			want: "cab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringShift(tt.args.s, tt.args.shift), "stringShift(%v, %v)", tt.args.s, tt.args.shift)
		})
	}
}
