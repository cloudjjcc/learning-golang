package nowcoder

import "testing"

func Test_numOfK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_01",
			args: args{
				arr: []int{1, 2, 2, 3},
				k:   2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfK(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("numOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
