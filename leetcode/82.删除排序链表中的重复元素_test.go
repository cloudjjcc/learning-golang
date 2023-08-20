package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_deleteDuplicates", args: args{BuildList([]int{1, 2, 3, 3, 4, 4, 5})}, want: []int{1, 2, 5},
		},
		{
			name: "Test_deleteDuplicates", args: args{BuildList([]int{1, 2, 2})}, want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ListToSlice(deleteDuplicates(tt.args.head)), "deleteDuplicates(%v)", tt.args.head)
		})
	}
}
