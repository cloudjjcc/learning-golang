package nowcoder

import "testing"

func Test_cutRope(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{i: 8},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutRope(tt.args.i); got != tt.want {
				t.Errorf("cutRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
