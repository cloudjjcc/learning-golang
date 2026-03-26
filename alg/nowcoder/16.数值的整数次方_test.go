package nowcoder

import "testing"

func Test_power(t *testing.T) {
	type args struct {
		base float64
		exp  int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test_01",
			args: args{
				base: 10,
				exp:  5,
			},
			want: 1e5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power(tt.args.base, tt.args.exp); got != tt.want {
				t.Errorf("power() = %v, want %v", got, tt.want)
			}
		})
	}
}
