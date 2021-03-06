package nowcoder

import (
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_01",
			args: args{str: "abcdefg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permutation(tt.args.str)
			t.Log(got)
		})
	}
}
