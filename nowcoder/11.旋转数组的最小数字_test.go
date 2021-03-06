package nowcoder

import "testing"

func Test_minNumberInRotateArray(t *testing.T) {
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
			args: args{arr: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "test_02",
			args: args{arr: []int{2, 1, 2, 2, 2}},
			want: 1,
		},
		{
			name: "test_03",
			args: args{arr: []int{1, 1, 1, 1, 1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberInRotateArray(tt.args.arr); got != tt.want {
				t.Errorf("minNumberInRotateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
