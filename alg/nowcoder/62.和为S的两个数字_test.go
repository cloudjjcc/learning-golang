package nowcoder

import (
	"reflect"
	"testing"
)

func Test_findNumsWithSum(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_01",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
				sum: 10,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumsWithSum(tt.args.arr, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumsWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
