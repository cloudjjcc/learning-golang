package nowcoder

import (
	"reflect"
	"testing"
)

func Test_findOnceNum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_01",
			args: args{arr: []int{1, 2, 3, 3, 4, 5, 5, 6, 1, 2}},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOnceNum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOnceNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
