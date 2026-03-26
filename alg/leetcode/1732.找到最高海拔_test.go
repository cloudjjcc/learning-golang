package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestAltitude(t *testing.T) {
	type args struct {
		gain []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_largestAltitude",
			args: args{gain: []int{-4, -3, -2, -1, 4, 3, 2}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, largestAltitude(tt.args.gain), "largestAltitude(%v)", tt.args.gain)
		})
	}
}
