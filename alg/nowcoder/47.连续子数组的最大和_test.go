package nowcoder

import "testing"

func Test_findGreatestSumOfSubArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{arr: []int{6, -3, -2, 7, -15, 1, 2, 2}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGreatestSumOfSubArray(tt.args.arr); got != tt.want {
				t.Errorf("findGreatestSumOfSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
