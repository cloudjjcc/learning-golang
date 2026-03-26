package nowcoder

import "testing"

func Test_getUglyNum(t *testing.T) {
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
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUglyNum(tt.args.n); got != tt.want {
				t.Errorf("getUglyNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
