package nowcoder

import "testing"

func TestPrintRepeatNum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testprintrepeatnum",
			args: args{nums: []int{4, 1, 1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintRepeatNum(tt.args.nums)
			PrintRepeatNum2(tt.args.nums)
		})
	}
}
