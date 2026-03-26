package datastructures

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	tests := []interface{}{
		1, "hello world", struct {
			name string
			age  int
		}{
			"xiao ming", 18,
		},
		[]int{3, 4},
	}
	list := NewSingleList()
	for _, tt := range tests {
		list.PushBack(tt)
	}
}

func TestLinkedList_Delete(t *testing.T) {

}

func TestLinkedList_Get(t *testing.T) {

}

func TestLinkedList_Insert(t *testing.T) {

}

func TestLinkedList_Size(t *testing.T) {

}
