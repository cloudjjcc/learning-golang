package datastructures

import (
	"fmt"
	"testing"
)

func TestBuildTreeFromArray(t *testing.T) {
	type args struct {
		arr []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0},
		{"test2", args{[]interface{}{0, 1, 2, 3, 4, 5}}, 0},
		{"test3", args{[]interface{}{0, 1, 2, 3, 4, 5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *TestTreeNode
			if got = BuildTreeFromArray(tt.args.arr); got.Value != tt.want {
				t.Errorf("BuildTreeFromArray() = %v, want %v", got, tt.want)
			}
			fmt.Print(got.String())
		})
	}
}

func TestTestTreeNode_InOrderNotRecursive(t1 *testing.T) {
	tree := BuildTreeFromArray([]interface{}{4, 2, 6, nil, 3, 5, nil})
	tree.InOrderNotRecursive(func(node *TestTreeNode) {
		fmt.Println(node.Value)
	})
}

func TestTestTreeNode_PreOrderNotRecursive(t1 *testing.T) {
	tree := BuildTreeFromArray([]interface{}{4, 2, 6, nil, 3, 5, nil})
	tree.PreOrderNotRecursive(func(node *TestTreeNode) {
		fmt.Println(node.Value)
	})
}
