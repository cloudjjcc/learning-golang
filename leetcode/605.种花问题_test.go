package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_canPlaceFlowers",
			args: args{n: 1, flowerbed: []int{1, 0, 0, 0, 1}},
			want: true,
		},
		{
			name: "Test_canPlaceFlowers",
			args: args{n: 2, flowerbed: []int{1, 0, 0, 0, 1, 0, 0}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canPlaceFlowers(tt.args.flowerbed, tt.args.n), "canPlaceFlowers(%v, %v)", tt.args.flowerbed, tt.args.n)
		})
	}
}
