package nowcoder

import "testing"

func Test_isExist(t *testing.T) {
	type args struct {
		matrix [][]int
		i      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_01",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				i: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExist(tt.args.matrix, tt.args.i); got != tt.want {
				t.Errorf("isExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
