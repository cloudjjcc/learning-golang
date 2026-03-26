package nowcoder

import "testing"

func TestPrintMinNum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_01",
			args: args{arr: []int{3, 32, 321}},
			want: "321323",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintMinNum(tt.args.arr); got != tt.want {
				t.Errorf("PrintMinNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
