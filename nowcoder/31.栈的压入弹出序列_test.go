package nowcoder

import (
	"testing"
)

func Test_isPopOrder(t *testing.T) {
	type args struct {
		pushOrder []int
		popOrder  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_01",
			args: args{
				pushOrder: []int{1, 2, 3},
				popOrder:  []int{1, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPopOrder(tt.args.pushOrder, tt.args.popOrder); got != tt.want {
				t.Errorf("isPopOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
