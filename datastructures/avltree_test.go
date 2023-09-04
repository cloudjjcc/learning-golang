package datastructures

import "testing"

func TestAVLTree_Add(t1 *testing.T) {
	type args struct {
		keys []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{keys: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &AVLTree{}
			for _, v := range tt.args.keys {
				t.Add(v)
			}
		})
	}
}
