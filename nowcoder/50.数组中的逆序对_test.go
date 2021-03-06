package nowcoder

import "testing"

func Test_inversePairs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{arr: []int{1, 3, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inversePairs(tt.args.arr); got != tt.want {
				t.Errorf("inversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
