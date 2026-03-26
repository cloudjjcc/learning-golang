package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topView(t *testing.T) {
	tree := buildTree([]int{1, 2, 3, 5, 7, 8, 6}, []int{2, 1, 7, 5, 8, 3, 6})
	assert.Equal(t, []int{2, 1, 3, 6}, topView(tree))
}
