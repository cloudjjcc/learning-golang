package nowcoder

import "testing"

func Test_hasPath(t *testing.T) {
	type args struct {
		matrix [][]rune
		path   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_01",
			args: args{
				matrix: [][]rune{{'a', 'b', 't', 'g'}, {'c', 'f', 'c', 's'}, {'j', 'd', 'e', 'h'}},
				path:   "bfcedh",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.args.matrix, tt.args.path); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
