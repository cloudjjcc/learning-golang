package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test_findMedianSortedArrays",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "Test_findMedianSortedArrays",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMedianSortedArrays(tt.args.nums1, tt.args.nums2), "findMedianSortedArrays(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
