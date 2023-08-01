package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_asteroidCollision", args: args{asteroids: []int{10, 2, -5}},
			want: []int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, asteroidCollision(tt.args.asteroids), "asteroidCollision(%v)", tt.args.asteroids)
		})
	}
}
