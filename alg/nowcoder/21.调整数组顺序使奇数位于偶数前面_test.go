package nowcoder

import "testing"

func Test_reOrderEvenOdd(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-01",
			args: args{arr: []int{1, 2, 3, 4, 5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reOrderEvenOdd(tt.args.arr)
		})
	}
}
