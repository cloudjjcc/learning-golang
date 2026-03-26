package nowcoder

import "testing"

func Test_countWay(t *testing.T) {
	type args struct {
		matrix [][]int
		x      int
		y      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{
				matrix: [][]int{},
				x:      50,
				y:      50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWay(tt.args.matrix, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("countWay() = %v, want %v", got, tt.want)
			}
		})
	}
}
