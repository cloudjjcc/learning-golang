package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_removeStars",
			args: args{s: "leet**cod*e"},
			want: "lecoe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeStars(tt.args.s), "removeStars(%v)", tt.args.s)
		})
	}
}
