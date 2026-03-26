package nowcoder

import "testing"

func Test_jumpFloor(t *testing.T) {
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
			args: args{i: 9},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpFloor(tt.args.i); got != tt.want {
				t.Errorf("jumpFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
