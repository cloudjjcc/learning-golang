package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test_pathSum",
			args: args{
				root:      BuildTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}, 0),
				targetSum: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pathSum(tt.args.root, tt.args.targetSum), "pathSum(%v, %v)", tt.args.root, tt.args.targetSum)
		})
	}
}
