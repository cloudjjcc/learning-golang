package huawei

import (
	"fmt"
	"strings"
	"testing"
)

func Test_lcp(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{strs: []string{
				"flower", "flow", "flight",
			}},
			want: "fl",
		},
		{
			name: "",
			args: args{strs: []string{
				"dog", "racecar", "car",
			}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcp(tt.args.strs); got != tt.want {
				t.Errorf("lcp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mainlcp(t *testing.T) {
	a := "[\"abc\",\"dsd\",\"edd\"]"
	a = strings.ReplaceAll(a, "[", "")
	a = strings.ReplaceAll(a, "]", "")
	a = strings.ReplaceAll(a, "\"", "")
	fmt.Println(a)
}
