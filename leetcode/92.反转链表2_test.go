package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_reverseBetween",
			args: args{
				head:  BuildList([]int{3, 5}),
				left:  1,
				right: 2,
			},
			want: []int{5, 3},
		},
		{
			name: "Test_reverseBetween",
			args: args{
				head:  BuildList([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ListToSlice(reverseBetween(tt.args.head, tt.args.left, tt.args.right)), "reverseBetween(%v, %v, %v)", tt.args.head, tt.args.left, tt.args.right)
		})
	}
}
