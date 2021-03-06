package nowcoder

import "testing"

func Test_getResultRect(t *testing.T) {
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
			args: args{n: 40},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResultRect(tt.args.n); got != tt.want {
				t.Errorf("getResultRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
