package nowcoder

import "testing"

func Test_numberOf1BetweenN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{n: 10},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOf1BetweenN(tt.args.n); got != tt.want {
				t.Errorf("numberOf1BetweenN() = %v, want %v", got, tt.want)
			}
		})
	}
}
