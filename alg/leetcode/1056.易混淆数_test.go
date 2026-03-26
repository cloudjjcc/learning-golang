package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_confusingNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_confusingNumber",
			args: args{n: 25},
			want: false,
		},
		{
			name: "Test_confusingNumber",
			args: args{n: 6},
			want: true,
		},
		{
			name: "Test_confusingNumber",
			args: args{n: 89},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, confusingNumber(tt.args.n), "confusingNumber(%v)", tt.args.n)
		})
	}
}
