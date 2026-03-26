package nowcoder

import "testing"

func Test_minCoinNum(t *testing.T) {
	type args struct {
		coins []int
		aim   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{
				coins: []int{2, 5, 10},
				aim:   7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCoinNum(tt.args.coins, tt.args.aim); got != tt.want {
				t.Errorf("minCoinNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
