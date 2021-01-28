package datastructures

import (
	"container/heap"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/emirpasic/gods/lists/singlylinkedlist"
)

var (
	singlyList singlylinkedlist.List
	doublyList doublylinkedlist.List
	arrayList  arraylist.List
	pqueue     heap.Interface
)
