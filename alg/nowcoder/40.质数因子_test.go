package nowcoder

import (
	"reflect"
	"testing"
)

func Test_getPrimer(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_01",
			args: args{n: 180},
			want: []int{2, 2, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPrimer(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got(%v) not equal want(%v)", got, tt.want)
			}
		})
	}
}
