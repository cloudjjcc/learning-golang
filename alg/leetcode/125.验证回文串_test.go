package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isPalindrome",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPalindrome(tt.args.s), "isPalindrome(%v)", tt.args.s)
		})
	}
}
