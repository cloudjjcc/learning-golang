package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_maxVowels",
			args: args{
				s: "abciiidef",
				k: 3,
			},
			want: 3,
		},
		{
			name: "Test_maxVowels",
			args: args{
				s: "weallloveyou",
				k: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxVowels(tt.args.s, tt.args.k), "maxVowels(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
