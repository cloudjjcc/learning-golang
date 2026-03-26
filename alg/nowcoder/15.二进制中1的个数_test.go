package nowcoder

import "testing"

func Test_numsOf1(t *testing.T) {
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
			args: args{i: 0x1010101},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numsOf1(tt.args.i); got != tt.want {
				t.Errorf("numsOf1() = %v, want %v", got, tt.want)
			}
		})
	}
}
