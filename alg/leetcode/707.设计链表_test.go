package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMyLinkedList(t *testing.T) {
	list := NewMyLinkedList()
	list.AddAtHead(1)
	fmt.Println(list)
	list.AddAtTail(3)
	fmt.Println(list)
	list.AddAtIndex(1, 2)
	fmt.Println(list)
	assert.Equal(t, 2, list.Get(1))
	list.DeleteAtIndex(1)
	fmt.Println(list)
	assert.Equal(t, 3, list.Get(1))
}
