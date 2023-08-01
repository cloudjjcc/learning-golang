package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_decodeString",
			args: args{s: "3[a]2[bc]"},
			want: "aaabcbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, decodeString(tt.args.s), "decodeString(%v)", tt.args.s)
		})
	}
}
