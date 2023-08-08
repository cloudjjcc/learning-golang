package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestSort",
			args: args{arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "TestSort",
			args: args{arr: []int{9, 8, 6, 6, 7, 4, 1, 2, 10}},
			want: []int{1, 2, 4, 6, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//BubbleSort(tt.args.arr)
			//InsertionSort(tt.args.arr)
			//SelectionSort(tt.args.arr)
			//MergeSort(tt.args.arr)
			//HeapSort(tt.args.arr)
			QuickSort(tt.args.arr)
			assert.Equal(t, tt.want, tt.args.arr)
		})
	}
}
func BenchmarkSort(b *testing.B) {
	buildArrFn := func(l int) []int {
		arr := make([]int, l)
		for i := 0; i < l; i++ {
			arr[i] = int(rand.Int63n(int64(l)))
		}
		return arr
	}
	sortFn := QuickSort
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortFn(buildArrFn(100000))
	}
	b.ReportAllocs()
}
