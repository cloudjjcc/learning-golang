package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_longestPalindrome2",
			args: args{s: "xaabacxcabaaxcabaax"},
			want: "xaabacxcabaax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPalindrome2(tt.args.s), "longestPalindrome2(%v)", tt.args.s)
		})
	}
}
