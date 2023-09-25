package datastructures

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipList(t *testing.T) {
	list := NewSkipList()
	list.Put(1, 1)
	list.Put(2, 2)
	assert.Equal(t, 1, list.Get(1))
	list.Del(1)
	assert.Equal(t, nil, list.Get(1))
	list.Put(3, 3)
	assert.Equal(t, 3, list.Get(3))
	for i := 0; i < 10000; i++ {
		list.Put(i, i)
	}
	fmt.Println(list.length, list.level)
}
