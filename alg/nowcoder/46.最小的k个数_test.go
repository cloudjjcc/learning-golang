package nowcoder

import (
	"reflect"
	"testing"
)

func Test_getLeastNums(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_01",
			args: args{
				arr: []int{4, 5, 1, 6, 2, 7, 3, 8},
				k:   4,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNums(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
