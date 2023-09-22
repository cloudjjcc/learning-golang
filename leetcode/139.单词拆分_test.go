package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_wordBreak",
			args: args{s: "leetcode", wordDict: []string{"leet", "code"}},
			want: true,
		},
		{
			name: "Test_wordBreak",
			args: args{s: "aaaaaaa", wordDict: []string{"aaaa", "aaa"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wordBreak(tt.args.s, tt.args.wordDict), "wordBreak(%v, %v)", tt.args.s, tt.args.wordDict)
		})
	}
}
