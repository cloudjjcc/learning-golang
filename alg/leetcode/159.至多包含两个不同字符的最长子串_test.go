package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_lengthOfLongestSubstringTwoDistinct",
			args: args{s: "eceba"},
			want: 3,
		},
		{
			name: "Test_lengthOfLongestSubstringTwoDistinct",
			args: args{s: "ccaabbb"},
			want: 5,
		},
		{
			name: "Test_lengthOfLongestSubstringTwoDistinct",
			args: args{s: "abaccc"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLongestSubstringTwoDistinct(tt.args.s), "lengthOfLongestSubstringTwoDistinct(%v)", tt.args.s)
		})
	}
}
