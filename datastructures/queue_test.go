package datastructures

import (
	"container/list"
	"reflect"
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		list list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"test1", fields{}, 1},
		{"test2", fields{}, 2},
		{"test3", fields{}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Empty(t *testing.T) {
	type fields struct {
		list list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		list list.List
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test1", fields{list.List{}}, args{1}},
		{"test2", fields{list.List{}}, args{'2'}},
		{"test3", fields{list.List{}}, args{struct {
		}{}}},
		{"test4", fields{list.List{}}, args{new(Queue)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			q.Enqueue(tt.args.data)
		})
	}
}

func TestQueue_Front(t *testing.T) {
	type fields struct {
		list list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	type fields struct {
		list list.List
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
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
