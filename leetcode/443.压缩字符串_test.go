package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_compress",
			args: args{chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}},
			want: 6,
		},
		{
			name: "Test_compress",
			args: args{chars: []byte{'a', 'b', 'c'}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, compress(tt.args.chars), "compress(%v)", tt.args.chars)
		})
	}
}
