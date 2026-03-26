package leetcode

import "testing"

func Test_sequencePrint(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_sequencePrint",
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sequencePrint(tt.args.n)
		})
	}
}
