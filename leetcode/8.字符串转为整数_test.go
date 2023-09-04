package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_myAtoi",
			args: args{s: "-21"},
			want: -21,
		},
		{
			name: "Test_myAtoi",
			args: args{s: "   2132000"},
			want: 2132000,
		},
		{
			name: "Test_myAtoi",
			args: args{s: "   4193 with words"},
			want: 4193,
		},
		{
			name: "Test_myAtoi",
			args: args{s: "words and 987"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, myAtoi(tt.args.s), "myAtoi(%v)", tt.args.s)
		})
	}
}
