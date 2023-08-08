package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_sortList",
			args: args{head: []int{4, 2, 1, 3}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test_sortList",
			args: args{head: []int{-1, 5, 3, 4, 0}},
			want: []int{-1, 0, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ListToSlice(sortList(BuildList(tt.args.head))), "sortList(%v)", tt.args.head)
		})
	}
}
