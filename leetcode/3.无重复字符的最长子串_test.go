package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_lengthOfLongestSubstring",
			args: args{s: "abba"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLongestSubstring(tt.args.s), "lengthOfLongestSubstring(%v)", tt.args.s)
		})
	}
}
