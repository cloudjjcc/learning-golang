package huawei

import "testing"

func Test_alive(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{arr: []int{5, 10, 8, -8, -5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alive(tt.args.arr); got != tt.want {
				t.Errorf("alive() = %v, want %v", got, tt.want)
			}
		})
	}
}
