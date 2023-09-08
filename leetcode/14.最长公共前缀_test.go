package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_longestCommonPrefix",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestCommonPrefix(tt.args.strs), "longestCommonPrefix(%v)", tt.args.strs)
		})
	}
}
