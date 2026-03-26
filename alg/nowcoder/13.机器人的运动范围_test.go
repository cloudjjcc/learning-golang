package nowcoder

import "testing"

func Test_moveCount(t *testing.T) {
	type args struct {
		k    int
		rows int
		cols int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{
				k:    15,
				rows: 100,
				cols: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveCount(tt.args.k, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("moveCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
