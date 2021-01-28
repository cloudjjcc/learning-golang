package datastructures

import "testing"

func TestHeap_Pop(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test1", fields{[]int{1, 2, 3, -1, 88, 0}}, 88},
		{"test2", fields{[]int{4, 2, 3, 0, 0, 0, 1, 2, 3}}, 4},
		{"test3", fields{[]int{1, 2, 3, 4, 2, 3, 9, 1}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewHeap(tt.fields.data)
			if got := pq.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Push(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		ele int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"test1", fields{data: []int{1, 2, 3}}, args{ele: 5}, 5},
		{"test2", fields{data: []int{2, 2, 3}}, args{ele: 0}, 3},
		{"test3", fields{data: []int{1, 2, 3}}, args{ele: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewHeap(tt.fields.data)
			pq.Push(tt.args.ele)
			if got := pq.Pop(); tt.want != got {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Size(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &Heap{
				data: tt.fields.data,
			}
			if got := pq.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
