package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSkiplist(t *testing.T) {
	skiplist := NewSkiplist()
	skiplist.Add(6)
	skiplist.Add(2)
	skiplist.Add(4)
	skiplist.Add(4)
	skiplist.Add(1)
	skiplist.Add(1)
	skiplist.Add(1)
	skiplist.Add(1)
	skiplist.Add(1)
	assert.True(t, skiplist.Erase(1))
	skiplist.Add(9)
	skiplist.Add(9)
	skiplist.Add(9)
	skiplist.Add(9)
	assert.True(t, skiplist.Search(9))
	skiplist.Add(9)
	skiplist.Add(9)
	skiplist.Add(9)
	skiplist.Add(9)
	assert.True(t, skiplist.Search(4))
	assert.True(t, skiplist.Search(1))
	assert.True(t, skiplist.Erase(4))
	assert.True(t, skiplist.Search(4))
	assert.False(t, skiplist.Erase(0))
	skiplist.Add(0)
	assert.True(t, skiplist.Search(0))
	assert.True(t, skiplist.Erase(0))
	assert.True(t, skiplist.Erase(9))
	assert.True(t, skiplist.Erase(2))
	assert.True(t, skiplist.Erase(6))
	skiplist.Add(1)
	skiplist.Add(2)
	skiplist.Add(3)
	assert.True(t, skiplist.Search(3))
}
