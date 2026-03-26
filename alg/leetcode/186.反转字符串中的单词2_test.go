package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test_reverseWords",
			args: args{s: []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}},
			want: []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseWords(tt.args.s)
			assert.Equal(t, tt.want, tt.args.s)
		})
	}
}
