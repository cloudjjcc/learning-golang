package datastructures

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Delete(t *testing.T) {
	type fields struct {
		root       *TreeNode
		Comparator Comparator
	}
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{
				root:       tt.fields.root,
				Comparator: tt.fields.Comparator,
			}
			if got := b.Delete(tt.args.node); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Find(t *testing.T) {
	type fields struct {
		root       *TreeNode
		Comparator Comparator
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{
				root:       tt.fields.root,
				Comparator: tt.fields.Comparator,
			}
			if got := b.Find(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Insert(t *testing.T) {
	type fields struct {
		root       *TreeNode
		Comparator Comparator
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{
				root:       tt.fields.root,
				Comparator: tt.fields.Comparator,
			}
			if got := b.Insert(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Root(t *testing.T) {
	type fields struct {
		root       *TreeNode
		Comparator Comparator
	}
	tests := []struct {
		name   string
		fields fields
		want   *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{
				root:       tt.fields.root,
				Comparator: tt.fields.Comparator,
			}
			if got := b.Root(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWith(t *testing.T) {
	type args struct {
		comparator Comparator
	}
	tests := []struct {
		name string
		args args
		want *BinaryTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWith(tt.args.comparator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMax(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMax(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
