package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseStr(tt.args.s, tt.args.k), "reverseStr(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
