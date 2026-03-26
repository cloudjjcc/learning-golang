package nowcoder

import "testing"

func Test_triangleCount(t *testing.T) {
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
			args: args{arr: []int{4, 6, 7, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleCount(tt.args.arr); got != tt.want {
				t.Errorf("triangleCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
